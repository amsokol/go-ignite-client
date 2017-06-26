package internal

import (
	"net/url"

	"github.com/blang/semver"
)

// SessionToken is session token type
type SessionToken string

// SessionTokenNil means session token is not provided
const SessionTokenNil = ""

// SuccessStatus is session token type
type SuccessStatus int64

// Version is response data from `version` command
type Version semver.Version

// Client is the interface providing the methods to execute REST API commands
type Client interface {
	Execute(v url.Values) ([]byte, error)
	IsFailed(successStatus SuccessStatus) bool
	GetError(successStatus SuccessStatus, error string) string
}

// ResponseHeader is common for all responses
type ResponseHeader interface {
	GetSuccessStatus() SuccessStatus
	GetError() string
	GetSessionToken() SessionToken
}

// ResponseSQLQueryClose is response for `qrycls` commands
// See https://apacheignite.readme.io/v1.3/docs/rest-api#section-sql-query-close for more details
type ResponseSQLQueryClose interface {
	ResponseHeader
	GetResponse() bool
}

// ResponseSQLQueryFetch is response for `qryfetch`, commands
// See https://apacheignite.readme.io/v1.3/docs/rest-api#section-sql-query-fetch for more details
type ResponseSQLQueryFetch interface {
	ResponseHeader
	GetResponse() SQLQueryResult
}

// ResponseSQLFieldsQueryExecute is response for `qryfldexe`, commands
// See https://apacheignite.readme.io/v1.3/docs/rest-api#section-sql-fields-query-execute for more details
type ResponseSQLFieldsQueryExecute interface {
	ResponseHeader
	GetResponse() SQLQueryResult
}

// SQLQueryResult is body of response for `qryfldexe`, command
// See https://apacheignite.readme.io/v1.3/docs/rest-api#section-sql-fields-query-execute for more details
type SQLQueryResult interface {
	GetItems() [][]interface{}
	GetLast() bool
	GetQueryID() int64
	GetFieldsMetadata() []FieldMetadata
}

// FieldMetadata is column list
// See https://apacheignite.readme.io/v1.3/docs/rest-api#section-sql-fields-query-execute for more details
type FieldMetadata interface {
	GetSchemaName() string
	GetTypeName() string
	GetFieldName() string
	GetFieldTypeName() string
}
