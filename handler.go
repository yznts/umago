package umago

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"
)

var (
	pixelPng = `iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNkYAAAAAYAAjCB0C8AAAAASUVORK5CYII=`
)

// NewHandler creates a new Umami event handler
// with a given configuration.
// Main purpose of this handler is to avoid including Umami client-side tracker.
// It can be used in multiple ways:
//   - you can send a plain event info with JS fetch (POST request with {"n": "name", "t": "title", "r": "referrer"})
//   - tracking pixel reference (f.e. <img src="/tracking/endpoint.png?r=referrer&t=title">)
func NewHandler(c Configuration) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// First, we need to unpack the data.
		// We are switching data unpacking behavior depending on the request type.
		var data struct {
			Name    string `json:"n"`
			Title   string `json:"t"`
			Referer string `json:"r"`
		}
		switch {
		// POST request means we're receiving JSON data
		case r.Method == http.MethodPost:
			if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			break
		// GET request with png extension means we're receiving data in query string
		case r.Method == http.MethodGet &&
			strings.HasSuffix(r.URL.Path, ".png"):
			// Set data from query string
			data.Name = r.URL.Query().Get("n")
			data.Title = r.URL.Query().Get("t")
			data.Referer = r.URL.Query().Get("r")
			break
		}
		// Close request body
		if err := r.Body.Close(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Compose event, client and send to the Umami server
		client := NewClientFromHttpRequest(r)
		event := NewEventFromHttpRequest(r)
		event.Name = data.Name
		event.Title = data.Title
		event.Referer = data.Referer
		event.Url = r.Referer()
		if err := Send(c, client, event); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Next, we need to return a response.
		// We are switching response behavior depending on the request type.
		switch {
		// We're not responding to POST requests
		case r.Method == http.MethodPost:
			break
		// GET request with png extension means we're returning an pixel image
		case r.Method == http.MethodGet && (strings.HasSuffix(r.URL.Path, ".png")):
			pixel, _ := base64.StdEncoding.DecodeString(pixelPng)
			w.Header().Set("Content-Type", "image/png")
			w.Write(pixel)
			break
		}
	}
}
