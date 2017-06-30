package v1

import (
	"net/url"

	core "github.com/amsokol/go-ignite-client/http"
)

// Client is the interface providing the methods to execute REST API commands
type Client interface {
	GetLog(path string, from *int, to *int) (log string, token string, err error)
	GetVersion() (version string, token string, err error)
	Decrement(cache string, key string, init *int64, delta int64) (value int64, nodeID string, token string, err error)
	Increment(cache string, key string, init *int64, delta int64) (value int64, nodeID string, token string, err error)
	GetCacheMetrics(cache string, destID string) (metrics core.CacheMetrics, nodeID string, token string, err error)
	SQLQueryClose(queryID int64) (ok bool, token string, err error)
	SQLQueryFetch(pageSize int64, queryID int64) (result core.SQLQueryResult, token string, err error)
	SQLFieldsQueryExecute(cache string, pageSize int64, query string, args url.Values) (result core.SQLQueryResult, token string, err error)
	Close() (err error)
}

// Client is providing the methods to execute REST API commands
type client struct {
	servers  []string
	username string
	password string
}

// NewClient returns new client
func NewClient(servers []string, username string, password string) Client {
	return &client{servers: servers, username: username, password: password}
}

// Close frees resources if it's necessary
func (c *client) Close() (err error) {
	return nil
}
