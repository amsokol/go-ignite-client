package v1

import (
	"net/url"

	"github.com/amsokol/go-ignite-client/http/v1/internal"
	"github.com/amsokol/go-ignite-client/http/v1/v10"
	"github.com/amsokol/go-ignite-client/http/v1/v13"
)

// SessionToken is session token type
type SessionToken internal.SessionToken

// SessionTokenNil means session token is not provided
const SessionTokenNil = internal.SessionTokenNil

// Version is response data from `version` command
type Version internal.Version

// ResponseSQLQueryClose is response for `qrycls` commands
// See https://apacheignite.readme.io/v1.3/docs/rest-api#section-sql-query-close for more details
type ResponseSQLQueryClose interface {
	internal.ResponseSQLQueryClose
}

// ResponseSQLQueryFetch is response for `qryfetch`, commands
// See https://apacheignite.readme.io/v1.3/docs/rest-api#section-sql-query-fetch for more details
type ResponseSQLQueryFetch interface {
	internal.ResponseSQLQueryFetch
}

// ResponseSQLFieldsQueryExecute is response for `qryfldexe`, commands
// See https://apacheignite.readme.io/v1.3/docs/rest-api#section-sql-fields-query-execute for more details
type ResponseSQLFieldsQueryExecute interface {
	internal.ResponseSQLFieldsQueryExecute
}

// SQLQueryResult is body of response for `qryfldexe`, command
// See https://apacheignite.readme.io/v1.3/docs/rest-api#section-sql-fields-query-execute for more details
type SQLQueryResult interface {
	internal.SQLQueryResult
}

// FieldMetadata is column list
// See https://apacheignite.readme.io/v1.3/docs/rest-api#section-sql-fields-query-execute for more details
type FieldMetadata interface {
	internal.FieldMetadata
}

// Client is the interface providing the methods to execute REST API commands
type Client interface {
	GetVersion() (Version, SessionToken, error)
	SQLQueryClose(queryID string) (bool, SessionToken, error)
	SQLQueryFetch(pageSize int64, queryID string) (SQLQueryResult, SessionToken, error)
	SQLFieldsQueryExecute(cacheName string, pageSize int64, query string, args url.Values) (SQLQueryResult, SessionToken, error)
}

// Client is providing the methods to execute REST API commands
type client struct {
	client internal.Client
}

// GetVersion command shows current Ignite version.
// See https://apacheignite.readme.io/v1.3/docs/rest-api#section-version for more details
func (c *client) GetVersion() (Version, SessionToken, error) {
	v, s, err := v10.Version(c.client)
	return Version(v), SessionToken(s), err
}

// SQLQueryClose closes query resources
// See https://apacheignite.readme.io/v1.3/docs/rest-api#section-sql-query-close for more details
func (c *client) SQLQueryClose(queryID string) (bool, SessionToken, error) {
	res, s, err := v13.SQLQueryClose(c.client, queryID)
	return res, SessionToken(s), err
}

// SQLQueryFetch gets next page for the query
// See https://apacheignite.readme.io/v1.3/docs/rest-api#section-sql-query-fetch for more details
func (c *client) SQLQueryFetch(pageSize int64, queryID string) (SQLQueryResult, SessionToken, error) {
	res, s, err := v13.SQLQueryFetch(c.client, pageSize, queryID)
	return res, SessionToken(s), err
}

// SQLFieldsQueryExecute runs sql fields query over cache.
// See https://apacheignite.readme.io/v1.3/docs/rest-api#section-sql-fields-query-execute for more details
func (c *client) SQLFieldsQueryExecute(cacheName string, pageSize int64, query string, args url.Values) (SQLQueryResult, SessionToken, error) {
	res, s, err := v13.SQLFieldsQueryExecute(c.client, cacheName, pageSize, query, args)
	return res, SessionToken(s), err
}

// Open returns client
func Open(servers []string, username string, password string) Client {
	return &client{client: internal.Open(servers, username, password)}
}
