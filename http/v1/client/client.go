package client

import (
	"net/url"

	"github.com/amsokol/go-ignite-client/http/types"
	"github.com/amsokol/go-ignite-client/http/v1/client/internal"
)

// Client is the interface providing the methods to execute REST API commands
type Client interface {
	Execute(v url.Values) ([]byte, error)
	IsFailed(successStatus types.SuccessStatus) bool
	GetError(successStatus types.SuccessStatus, error string) string
}

// Open returns client
func Open(servers []string, username string, password string) Client {
	return &internal.Client{Servers: servers, Username: username, Password: password}
}
