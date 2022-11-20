package api

import (
	"context"
	"net/http"
	"net/url"
	"time"
)

// Client is a generic http client for service to service communication
type Client struct {
	client  *http.Client
	baseURL url.URL
}

// New instantiates a new Client
func New(baseURL url.URL) *Client {
	if baseURL.Hostname() == "" {
		baseURL = url.URL{
			Scheme: "http",
			Host:   "localhost",
		}
	}

	return &Client{
		client: &http.Client{
			Timeout: 15 * time.Second,
		},
		baseURL: baseURL,
	}
}

// SpideyClient defines the methods that may be undertaken on spidey data
type SpideyClient interface {
	GetSpidey(ctx context.Context, id int) Spidey
}
