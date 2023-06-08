package umami

import (
	"encoding/json"
	"net/http"
)

func NewHandler(c Configuration) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Unpack data
		var data struct {
			Referer string `json:"r"`
		}
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		_ = r.Body.Close()
		// Compose event
		event := NewEventFromHttpRequest(r)
		event.Referer = data.Referer
		event.Url = r.Referer()
		// Compose client
		client := NewClientFromHttpRequest(r)
		// Send to the server
		_ = Send(c, client, event)
	}
}
