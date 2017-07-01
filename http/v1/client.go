package v1

import (
	"net/url"

	core "github.com/amsokol/go-ignite-client/http"

	"github.com/amsokol/go-ignite-client/http/v1/along"
	"github.com/amsokol/go-ignite-client/http/v1/exec"
	"github.com/amsokol/go-ignite-client/http/v1/server"
	"github.com/amsokol/go-ignite-client/http/v1/sql"
)

// Client is the interface providing the methods to execute REST API commands
type Client interface {
	// server
	GetVersion() (version string, token string, err error)
	GetLog(path string, from *int, to *int) (log string, token string, err error)

	// sql
	SQLFieldsQueryExecute(cache string, pageSize int64, query string, args url.Values) (result core.SQLQueryResult, token string, err error)
	SQLQueryFetch(pageSize int64, queryID int64) (result core.SQLQueryResult, token string, err error)
	SQLQueryClose(queryID int64) (ok bool, token string, err error)

	// along
	Decrement(cache string, key string, init *int64, delta int64) (value int64, nodeID string, token string, err error)
	Increment(cache string, key string, init *int64, delta int64) (value int64, nodeID string, token string, err error)

	Close() (err error)
}

// ClientImpl is providing the methods to execute REST API commands
type client struct {
	exec   exec.ExecuterImpl
	server server.Commands
	sql    sql.Commands
	along  along.Commands
}

// GetVersion command shows current Ignite version.
// See https://apacheignite.readme.io/v1.9/docs/rest-api#section-version for more details
func (c *client) GetVersion() (version string, token string, err error) {
	return c.server.GetVersion(&c.exec)
}

// GetLog command shows server logs
// See https://apacheignite.readme.io/v1.9/docs/rest-api#log for more details
func (c *client) GetLog(path string, from *int, to *int) (log string, token string, err error) {
	return c.server.GetLog(&c.exec, path, from, to)
}

// SQLFieldsQueryExecute runs sql fields query over cache.
// See https://apacheignite.readme.io/v1.9/docs/rest-api#section-sql-fields-query-execute for more details
func (c *client) SQLFieldsQueryExecute(cache string, pageSize int64, query string, args url.Values) (result core.SQLQueryResult, token string, err error) {
	return c.sql.SQLFieldsQueryExecute(&c.exec, cache, pageSize, query, args)
}

// SQLQueryFetch gets next page for the query
// See https://apacheignite.readme.io/v1.9/docs/rest-api#section-sql-query-fetch for more details
func (c *client) SQLQueryFetch(pageSize int64, queryID int64) (result core.SQLQueryResult, token string, err error) {
	return c.sql.SQLQueryFetch(&c.exec, pageSize, queryID)
}

// SQLQueryClose closes query resources
// See https://apacheignite.readme.io/v1.9/docs/rest-api#section-sql-query-close for more details
func (c *client) SQLQueryClose(queryID int64) (ok bool, token string, err error) {
	return c.sql.SQLQueryClose(&c.exec, queryID)
}

// Decrement command subtracts and gets current value of given atomic long
// See https://apacheignite.readme.io/v1.9/docs/rest-api#section-decrement for more details
func (c *client) Decrement(cache string, key string, init *int64, delta int64) (value int64, nodeID string, token string, err error) {
	return c.along.Decrement(&c.exec, cache, key, init, delta)
}

// Increment command adds and gets current value of given atomic long
// See https://apacheignite.readme.io/v1.9/docs/rest-api#section-increment for more details
func (c *client) Increment(cache string, key string, init *int64, delta int64) (value int64, nodeID string, token string, err error) {
	return c.along.Increment(&c.exec, cache, key, init, delta)
}

// NewClient returns new client
func NewClient(servers []string, username string, password string) Client {
	c := client{}
	c.exec.Servers = servers
	c.exec.Username = username
	c.exec.Password = password
	return &c
}

// Close frees resources if it's necessary
func (c *client) Close() (err error) {
	return nil
}
