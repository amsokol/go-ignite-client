package v2

import (
	"net/url"

	core "github.com/amsokol/go-ignite-client/http"
	"github.com/amsokol/go-ignite-client/http/v1"
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
	client v1.Client
}

// Log command shows server logs
// See https://apacheignite.readme.io/v1.0/docs/rest-api#log for more details
func (c *client) GetLog(path string, from *int, to *int) (log string, token string, err error) {
	return c.client.GetLog(path, from, to)
}

// Version command shows current Ignite version.
// See https://apacheignite.readme.io/v1.3/docs/rest-api#section-version for more details
func (c *client) GetVersion() (version string, token string, err error) {
	return c.client.GetVersion()
}

// Decrement command subtracts and gets current value of given atomic long
// See https://apacheignite.readme.io/v1.0/docs/rest-api#section-decrement for more details
func (c *client) Decrement(cache string, key string, init *int64, delta int64) (value int64, nodeID string, token string, err error) {
	return c.client.Decrement(cache, key, init, delta)
}

// Increment command adds and gets current value of given atomic long
// See https://apacheignite.readme.io/v1.0/docs/rest-api#section-increment for more details
func (c *client) Increment(cache string, key string, init *int64, delta int64) (value int64, nodeID string, token string, err error) {
	return c.client.Increment(cache, key, init, delta)
}

// CacheMetrics shows metrics for Ignite cache
// See https://apacheignite.readme.io/v1.0/docs/rest-api#section-cache-metrics for more details
func (c *client) GetCacheMetrics(cache string, destID string) (metrics core.CacheMetrics, nodeID string, token string, err error) {
	return c.client.GetCacheMetrics(cache, destID)
}

// SQLQueryClose closes query resources
// See https://apacheignite.readme.io/v1.3/docs/rest-api#section-sql-query-close for more details
func (c *client) SQLQueryClose(queryID int64) (ok bool, token string, err error) {
	return c.client.SQLQueryClose(queryID)
}

// SQLQueryFetch gets next page for the query
// See https://apacheignite.readme.io/v1.3/docs/rest-api#section-sql-query-fetch for more details
func (c *client) SQLQueryFetch(pageSize int64, queryID int64) (result core.SQLQueryResult, token string, err error) {
	return c.client.SQLQueryFetch(pageSize, queryID)
}

// SQLFieldsQueryExecute runs sql fields query over cache.
// See https://apacheignite.readme.io/v1.3/docs/rest-api#section-sql-fields-query-execute for more details
func (c *client) SQLFieldsQueryExecute(cache string, pageSize int64, query string, args url.Values) (result core.SQLQueryResult, token string, err error) {
	return c.client.SQLFieldsQueryExecute(cache, pageSize, query, args)
}

func (c *client) Close() (err error) {
	return c.Close()
}

// NewClient returns new client
func NewClient(servers []string, username string, password string) Client {
	return v1.NewClient(servers, username, password)
}
