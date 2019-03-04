package glo

import "net/http"

// Glo API object
type Glo struct {
	token   string
	client  *http.Client
	BaseURI string
}

// NewClient Glo API Client
func NewClient(token string) *Glo {
	return &Glo{
		client:  &http.Client{},
		token:   token,
		BaseURI: "https://gloapi.gitkraken.com/v1/glo",
	}
}
