package umami

import "net/http"

func NewHandler(c Configuration) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Compose client, event and send to the server
		_ = Send(c, NewClientFromHttpRequest(r), NewEventFromHttpRequest(r))
	}
}
