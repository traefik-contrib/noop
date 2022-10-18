package noop

import (
	"context"
	"errors"
	"net/http"
)

// Config the plugin configuration.
type Config struct {
	ResponseCode int `json:"responseCode"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		ResponseCode: http.StatusTeapot,
	}
}

// Noop a Noop plugin.
type Noop struct {
	next         http.Handler
	responseCode int
	name         string
}

// New created a new Noop plugin.
func New(_ context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	if config.ResponseCode < 100 || config.ResponseCode > 999 {
		return nil, errors.New("invalid response code")
	}

	return &Noop{
		responseCode: config.ResponseCode,
		next:         next,
		name:         name,
	}, nil
}

func (n *Noop) ServeHTTP(rw http.ResponseWriter, _ *http.Request) {
	rw.WriteHeader(n.responseCode)
}
