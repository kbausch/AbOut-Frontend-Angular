package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"gitlab.cs.mtech.edu/ESOF326/S20/AbOut/backend/middleware"

	"net/http"

	"gitlab.cs.mtech.edu/ESOF326/S20/AbOut/backend/server"
)

// Define Server timeout constants.
const (
	RWTimeout        = time.Second * 15
	IdleTimeout      = time.Second * 60
	ShutdownDeadline = time.Second * 15
)

var supportedAuthTypes = map[string]bool{"jwt": true, "cas": true}

func parseFlags() {
	// Check for options passed in from the command line.
	// auth is the type of authentication to use. jwt and cas are supported.
	authType := flag.CommandLine.String("auth", "jwt", "Specify the authentication method")

	// secret is the secret key for hashing jwt tokens. It can be changed from the default.
	secret := flag.CommandLine.String("secret", "̣㕋⬏ꖼ荾٥睻ꚏ媳⣦줨ꤏ蒲퉐⺿폻媒툆䟯隊듃睢㴢㴮ⅲԮ伈샍髵某肣", "Specify the secret string for encoding jwt tokens.")

	// Parse command line options.
	flag.Parse()

	if supportedAuthTypes[*authType] != true {
		fmt.Printf("Unsupported authentication type: %v\n", *authType)
		os.Exit(0)
	}
	os.Setenv("authType", *authType)
	log.Printf("Authentication type: %v\n", *authType)

	os.Setenv("secretKey", *secret)
}

func main() {
	// Process command line flags.
	parseFlags()

	// Grab the port to serve on from the OS variables or use a default.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := server.NewRouter()

	// Application of the CORs middleware
	h := middleware.CreateCorsMiddleware().Handler(r)

	// Application of the CAS middleware.
	if os.Getenv("authType") == "cas" {
		h = middleware.CreateCasClient().Handler(r)
	}

	// Create a server with timeouts and our router.
	srv := &http.Server{
		Addr: ":" + port,

		// It is good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: RWTimeout,
		ReadTimeout:  RWTimeout,
		IdleTimeout:  IdleTimeout,
		Handler:      h,
	}

	// Open a separate thread to handle requests so we can handle graceful
	// shutdown in the main thread.
	go func() {
		log.Printf("Listening on Port: %s\n", port)
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	// Create a channel waiting for a Operating System interrupt.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block this thread until we receive a signal from the os to shutdown.
	<-c

	// Create deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), ShutdownDeadline)
	defer cancel()

	// Shudown doesn't block if there are no connections, but will otherwise
	// wait until the timeout before closing connections and tearing down.
	srv.Shutdown(ctx)

	log.Println("shutting down")
	os.Exit(0)
}
