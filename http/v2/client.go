package v2

import (
	"github.com/amsokol/go-ignite-client/http/v1"
)

// Client is the interface providing the methods to execute REST API commands
type Client interface {
	v1.Client
}

// NewClient returns new client
func NewClient(servers []string, username string, password string) Client {
	return v1.NewClient(servers, username, password)
}
