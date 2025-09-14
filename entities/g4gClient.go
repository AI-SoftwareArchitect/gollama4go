package entities

import "net/http"

type G4GClient struct {
	BaseURL  string
	Http     *http.Client
	MaxQueue int
}
