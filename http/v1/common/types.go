package common

import (
	"net/url"

	"github.com/blang/semver"
)

// SessionToken is session token type
type SessionToken string

// SessionTokenNil means session token is not provided
// TODO: move to internal
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
