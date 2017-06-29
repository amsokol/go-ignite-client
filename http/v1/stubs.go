package v1

import (
	"net/url"

	"github.com/amsokol/go-ignite-client/http/types"
	v1 "github.com/amsokol/go-ignite-client/http/v1/client"
	"github.com/amsokol/go-ignite-client/http/v1/v10"
	"github.com/amsokol/go-ignite-client/http/v1/v13"
)

// Client is the interface providing the methods to execute REST API commands
type Client interface {
	Log(path string, from int, to int) (string, types.SessionToken, error)
	Version() (types.Version, types.SessionToken, error)
	SQLQueryClose(queryID string) (bool, types.SessionToken, error)
	SQLQueryFetch(pageSize int64, queryID string) (*types.SQLQueryResult, types.SessionToken, error)
	SQLFieldsQueryExecute(cacheName string, pageSize int64, query string, args url.Values) (*types.SQLQueryResult, types.SessionToken, error)
}

// Client is providing the methods to execute REST API commands
type client struct {
	client v1.Client
}

// Log command shows server logs
// See https://apacheignite.readme.io/v1.0/docs/rest-api#log for more details
func (c *client) Log(path string, from int, to int) (string, types.SessionToken, error) {
	return v10.Log(c.client, path, from, to)
}

// Version command shows current Ignite version.
// See https://apacheignite.readme.io/v1.3/docs/rest-api#section-version for more details
func (c *client) Version() (types.Version, types.SessionToken, error) {
	return v10.Version(c.client)
}

// SQLQueryClose closes query resources
// See https://apacheignite.readme.io/v1.3/docs/rest-api#section-sql-query-close for more details
func (c *client) SQLQueryClose(queryID string) (bool, types.SessionToken, error) {
	return v13.SQLQueryClose(c.client, queryID)
}

// SQLQueryFetch gets next page for the query
// See https://apacheignite.readme.io/v1.3/docs/rest-api#section-sql-query-fetch for more details
func (c *client) SQLQueryFetch(pageSize int64, queryID string) (*types.SQLQueryResult, types.SessionToken, error) {
	return v13.SQLQueryFetch(c.client, pageSize, queryID)
}

// SQLFieldsQueryExecute runs sql fields query over cache.
// See https://apacheignite.readme.io/v1.3/docs/rest-api#section-sql-fields-query-execute for more details
func (c *client) SQLFieldsQueryExecute(cacheName string, pageSize int64, query string, args url.Values) (*types.SQLQueryResult, types.SessionToken, error) {
	return v13.SQLFieldsQueryExecute(c.client, cacheName, pageSize, query, args)
}

// Open returns client
func Open(servers []string, username string, password string) Client {
	return &client{client: v1.Open(servers, username, password)}
}
