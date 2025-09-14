package entities

import (
	"net/http"
	"time"
)

func NewG4GClient(baseURL string, maxQueue int) *G4GClient {
	return &G4GClient{
		BaseURL: baseURL,
		Http: &http.Client{
			Timeout: 10 * time.Second,
		},
		MaxQueue: maxQueue,
	}
}
