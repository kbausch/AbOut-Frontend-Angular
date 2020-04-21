package middleware

import (
	"log"
	"net/url"

	"gopkg.in/Dolphindalt/cas.v4"
)

// casURL is the URL for all CAS requests.
const casURL = "https://mtlbsso.mtech.edu:443/idp/profile/cas"

// CreateCasClient creates middleware for CAS authentication.
func CreateCasClient() *cas.Client {
	url, _ := url.Parse(casURL)
	log.Printf("CAS URL: %v\n", casURL)
	client := cas.NewClient(&cas.Options{URL: url, SendService: true})
	return client
}
