# Middleware for Permissions

This is a small document detailing the plans of the implementation for the middleware that will handle permissions and authentication in the AbOut system.

## What is Middleware?

Essentially, middleware performs some specific function on the HTTP request or response at a specific stage in the HTTP pipeline before or after the user defined endpoint. Middleware is used to implement logging, handling authentication, and gzip compression without having many code contact points.

In Go, middleware is simply a function with the prototype shown below.

```go
// Middleware defines the function signature for a middleware.
type Middleware func(http.HandlerFunc) http.HanderFunc
```

In Go, middleware is really just a function that takes the endpoint or another middleware as a parameter, applies middleware operations to itself, and then calls the next http handler in the chain. A convenient function is given below for chaining multiple middlewares to an endpoint handler in a single call.

```go
// Chain applies middlewares to an http.Handler.
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}
```

Since each middleware is defined separately from each endpoint and middlewares can be attached to each endpoint, the code is much more modular. Each middleware could be thought of as a module that can be applied to an arbitrary endpoint with minimal boilerplate code.

## Authentication

The authentication process begins with the user attempting to login to some web service with a username and password. If the credentials are correct, the user will receive a token as proof of their authentication. The token is encoded with information about the user such as their username. In subsequent http requests, the token will be sent in the request in a header in the form shown below.

```
Authorization: Bearer <token>
```

A middleware can be created to intercept this header from the request and check if the user is authenticated. If they are, the information will be attached to the request and the endpoint will decide what to do with that information. An endpoint could refuse all unauthenticated users or it could only allow access for authenticated users who hold certain permissions. But how will the endpoint know what permissions the user has without putting permission grabbing code in each endpoint? Middleware is the solution.

## Authentication Type

As for what authentication method will be supported, this application will provide an interface for an arbitrary form of authentication. The interface is defined as follows.

```go
type authenticationType interface {
    isAuthenticated(string) bool
    authenticateUser(string, string) bool
}
```

A generic interface is provided so that we may support both JWT based authentication and CAS based authentication. Given that dealing with CAS has been difficult in the past, authentication may be done with JWT temporarily. The interface will allow both to be supported.

## Getting Permissions

Now that we have an interface to handle authentication, a middleware must be designed to handle managing user permissions.

First, the prototype function for fetching the permission data from the database is given. Keep in mind the interface{} type is meant to represent a generic JSON value.

```go
func fetchPermissions(username string) interface{}
```

These permissions in the interface function will then be appended to the request body which is also in JSON format. It is unclear how this will look exactly at this moment in time but the decision on this will be made during sprint 2 after the database team completes their stored procedure for permissions in sprint 1.

Now that the middleware has appended the permissions to the request body, the endpoints will be able to process them. Helper functions for doing this will be designed during sprint 2.