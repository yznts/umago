package umami

import (
	"net/http"
	"strings"
)

type Event struct {
	Hostname string `json:"hostname"`
	Language string `json:"language"`
	Referer  string `json:"referrer"`
	Screen   string `json:"screen"`
	Title    string `json:"title"`
	Url      string `json:"url"`
	Website  string `json:"website"`
	Name     string `json:"name"`
}

func NewEventFromHttpRequest(r *http.Request) Event {
	// Resolve language
	language := r.Header.Get("Accept-Language")
	if language == "*" {
		language = ""
	}
	if language != "" {
		language = strings.Split(language, ",")[0]
	}
	// Compose event
	return Event{
		Hostname: r.Host,
		Language: language,
		Referer:  r.Referer(),
		Url:      r.URL.String(),
	}
}
