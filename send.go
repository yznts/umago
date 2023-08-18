package umago

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Send is a function to send an event to Umami
// with a given configuration and client.
func Send(conf Configuration, client Client, event Event) error {
	// Set event website if missing
	if event.Website == "" {
		event.Website = conf.Website
	}
	// Serialize and pack event
	eventPkg, err := json.Marshal(map[string]any{
		"type":    "event",
		"payload": event,
	})
	if err != nil {
		return fmt.Errorf("event serialization failed: %s", err.Error())
	}
	// Compose request
	req, err := http.NewRequest("POST", conf.Href+"/api/send", bytes.NewReader(eventPkg))
	if err != nil {
		return fmt.Errorf("request composition failed: %s", err.Error())
	}
	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", client.UserAgent)
	req.Header.Set("X-Client-IP", client.IP)
	// Send request
	res, err := http.DefaultClient.Do(req)
	// Check error
	if err != nil {
		return fmt.Errorf("request failed: %s", err.Error())
	}
	// Ensure body close
	defer res.Body.Close()
	// Check status
	if res.StatusCode != 200 {
		// Compose error with details
		body, _ := io.ReadAll(res.Body)
		return fmt.Errorf(`request failed: %s
			pkg: %s
			res: %s`,
			res.Status, string(eventPkg), string(body))
	}
	// Done
	return nil
}
