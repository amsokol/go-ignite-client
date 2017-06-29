package types

import (
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

// SQLQueryResult is body of response for `qryfetch`, command
// See https://apacheignite.readme.io/v1.3/docs/rest-api#section-sql-query-fetch for more details
type SQLQueryResult struct {
	Items          [][]interface{} `json:"items"`
	Last           bool            `json:"last"`
	QueryID        int64           `json:"queryId"`
	FieldsMetadata []FieldMetadata `json:"fieldsMetadata"`
}

// FieldMetadata is column list
// See https://apacheignite.readme.io/v1.3/docs/rest-api#section-sql-fields-query-execute for more details
type FieldMetadata struct {
	SchemaName    string `json:"schemaName"`
	TypeName      string `json:"typeName"`
	FieldName     string `json:"fieldName"`
	FieldTypeName string `json:"fieldTypeName"`
}

// CacheMetrics is the response for Cache metrics command
// See https://apacheignite.readme.io/v1.0/docs/rest-api#section-cache-metrics for more details
type CacheMetrics struct {
	CreateTime int64 `json:"createTime"`
	Hits       int64 `json:"hits"`
	Misses     int64 `json:"misses"`
	ReadTime   int64 `json:"readTime"`
	Reads      int64 `json:"reads"`
	WriteTime  int64 `json:"writeTime"`
	Writes     int64 `json:"writes"`
}
