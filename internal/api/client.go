package api

import (
	"net/http"
	"time"
)

// TODO : Add caching later if needed
type Client struct {
	httpClient http.Client
}

func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
