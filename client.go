package umago

import (
	"net/http"
	"strings"
)

type Client struct {
	IP        string
	UserAgent string
}

func NewClientFromHttpRequest(r *http.Request) Client {
	// Resolve ip
	ip := r.RemoteAddr
	{
		priority := []string{
			"X-Client-IP",
			"X-Forwarded-For",
			"CF-Connecting-IP",
			"Fastly-Client-IP",
			"True-Client-IP",
			"X-Real-IP",
			"X-Cluster-Client-IP",
			"X-Forwarded",
			"Forwarded-For",
			"Forwarded",
		}
		for _, header := range priority {
			if r.Header.Get(header) != "" {
				ip = strings.Split(r.Header.Get(header), ",")[0]
				break
			}
		}
	}
	// Compose client
	return Client{
		IP:        ip,
		UserAgent: r.UserAgent(),
	}
}
