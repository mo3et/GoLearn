// A middleware in itself simply takes a as one of its parameters
// warps it and returns a new for the server to call  http.HandlerFunc

// Example:
// func createNewMiddleware() Middleware {
// 	// Create a new Middleware
// 	middleware := func(next http.HandlerFunc) http.HandlerFunc {

// 		// Define the http.HandlerFunc which is called by the server eventually
// 		handler := func(w http.ResponseWriter, r *http.Request) {

// 			// ... do middleware things

// 			// Call the next middleware/handler in chain
// 			next(w, r)
// 		}

// 		// Return newly created handler
// 		return handler
// 	}

// 	// Return newly created middleware
// 	return middleware
// }

// ===========

// full example

// advanced-middleware.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

// Logging logs all requests with its path and the time it took to process
func Logging() Middleware {

	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {

			// Do middleware things
			start := time.Now()
			defer func() { log.Println(r.URL.Path, time.Since(start)) }()

			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}

// Method ensures that url can only be requested with a specific method, else returns a 400 Bad Request
func Method(m string) Middleware {

	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {

			// Do middleware things
			if r.Method != m {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}

// Chain applies middlewares to a http.HandlerFunc
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}
func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func main() {
	http.HandleFunc("/", Chain(Hello, Method("GET"), Logging())) //
	http.ListenAndServe(":8088", nil)
}
