package cwa

import (
	"net/http"
	"time"
)

// Client is the entry point for interacting with Taiwan CWA API.
type Client struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
}

// NewClient creates a new Client with the given API key and options.
func NewClient(apiKey string, opts ...Option) *Client {
	c := &Client{
		apiKey:  apiKey,
		baseURL: "https://opendata.cwa.gov.tw/api",
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}
