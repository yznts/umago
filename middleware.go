package umami

import (
	"net/http"
)

type Middleware func(next http.HandlerFunc) http.HandlerFunc

func NewMiddleware(c Configuration) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// Call next handler
			next(w, r)
			// Compose client, event and send to the server
			_ = Send(c, NewClientFromHttpRequest(r), NewEventFromHttpRequest(r))
		}
	}
}
