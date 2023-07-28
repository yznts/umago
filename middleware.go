package umago

import (
	"log"
	"net/http"
)

// Middleware is a function that takes a handler and returns a new handler.
type Middleware func(next http.HandlerFunc) http.HandlerFunc

// NewMiddleware returns a new middleware handler with given configuration.
// Please note that middleware can't handle client side URL changes
// or cached page loads.
// Before using it, please, make sure that middleware works for you.
func NewMiddleware(c Configuration) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// First, we are calling original handler
			// to not block the request processing.
			next(w, r)
			// Here we are extracting client info,
			// composing event info from request
			// and sending it to the server.
			err := Send(c, NewClientFromHttpRequest(r), NewEventFromHttpRequest(r))
			// If something went wrong, we are logging it.
			if err != nil {
				log.Printf("umami error: %v", err)
			}
		}
	}
}
