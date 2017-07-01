package v1

import (
	"net/url"

	core "github.com/amsokol/go-ignite-client/http"

	"github.com/amsokol/go-ignite-client/http/v1/along"
	"github.com/amsokol/go-ignite-client/http/v1/cache"
	"github.com/amsokol/go-ignite-client/http/v1/exec"
	"github.com/amsokol/go-ignite-client/http/v1/kvpairs"
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

	// kvpairs
	Add(cache string, key string, val string, destID string) (ok bool, nodeID string, token string, err error)
	Append(cache string, key string, val string, destID string) (ok bool, nodeID string, token string, err error)
	CompareAndSwap(cache string, key string, val string, val2 string, destID string) (ok bool, nodeID string, token string, err error)
	Prepend(cache string, key string, val string, destID string) (ok bool, nodeID string, token string, err error)
	Replace(cache string, key string, val string, val2 string, destID string) (ok bool, nodeID string, token string, err error)

	// cache
	GetOrCreateCache(cache string) (token string, err error)
	DestroyCache(cache string) (token string, err error)
	GetCacheMetrics(cache string, destID string) (metrics core.CacheMetrics, nodeID string, token string, err error)

	Close() (err error)
}

// ClientImpl is providing the methods to execute REST API commands
type client struct {
	exec    exec.ExecuterImpl
	server  server.Commands
	sql     sql.Commands
	along   along.Commands
	kvpairs kvpairs.Commands
	cache   cache.Commands
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

// Add command stores a given key-value pair in cache only if there isn't a previous mapping for it
// See https://apacheignite.readme.io/v1.9/docs/rest-api#section-add for for details
func (c *client) Add(cache string, key string, val string, destID string) (ok bool, nodeID string, token string, err error) {
	return c.kvpairs.Add(&c.exec, cache, key, val, destID)
}

// Append appends a line for value which is associated with key
// See https://apacheignite.readme.io/v1.9/docs/rest-api#section-append for more details
func (c *client) Append(cache string, key string, val string, destID string) (ok bool, nodeID string, token string, err error) {
	return c.kvpairs.Append(&c.exec, cache, key, val, destID)
}

// CompareAndSwap stores given key-value pair in cache only if the previous value is equal to the expected value passed in
// See https://apacheignite.readme.io/v1.9/docs/rest-api#section-compare-and-swap for details
func (c *client) CompareAndSwap(cache string, key string, val string, val2 string, destID string) (ok bool, nodeID string, token string, err error) {
	return c.kvpairs.CompareAndSwap(&c.exec, cache, key, val, val2, destID)
}

// Prepend prepends a line for value which is associated with key
// See https://apacheignite.readme.io/v1.9/docs/rest-api#section-prepend for more details
func (c *client) Prepend(cache string, key string, val string, destID string) (ok bool, nodeID string, token string, err error) {
	return c.kvpairs.Prepend(&c.exec, cache, key, val, destID)
}

// Replace stores a given key-value pair in cache only if there is a previous mapping for it
// See https://apacheignite.readme.io/v1.9/docs/rest-api#section-replace for more details
func (c *client) Replace(cache string, key string, val string, val2 string, destID string) (ok bool, nodeID string, token string, err error) {
	return c.kvpairs.Replace(&c.exec, cache, key, val, val2, destID)
}

// GetCacheMetrics shows metrics for Ignite cache
// See https://apacheignite.readme.io/v1.9/docs/rest-api#section-cache-metrics for more details
func (c *client) GetOrCreateCache(cache string) (token string, err error) {
	return c.cache.GetOrCreateCache(&c.exec, cache)
}

// DestroyCache destroys cache with given name
// See https://apacheignite.readme.io/v1.9/docs/rest-api#section-destroy-cache for more details
func (c *client) DestroyCache(cache string) (token string, err error) {
	return c.cache.DestroyCache(&c.exec, cache)
}

// GetOrCreateCache creates cache with given name if it does not exist
// See https://apacheignite.readme.io/v1.9/docs/rest-api#section-get-or-create-cache for more details
func (c *client) GetCacheMetrics(cache string, destID string) (metrics core.CacheMetrics, nodeID string, token string, err error) {
	return c.cache.GetCacheMetrics(&c.exec, cache, destID)
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
