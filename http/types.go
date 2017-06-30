package http

// SQLQueryResult is body of response for `qryfetch`, command
// See https://apacheignite.readme.io/v1.9/docs/rest-api#section-sql-query-fetch for more details
type SQLQueryResult struct {
	Items          [][]interface{} `json:"items"`
	Last           bool            `json:"last"`
	QueryID        int64           `json:"queryId"`
	FieldsMetadata []FieldMetadata `json:"fieldsMetadata"`
}

// FieldMetadata is column list
// See https://apacheignite.readme.io/v1.9/docs/rest-api#section-sql-fields-query-execute for more details
type FieldMetadata struct {
	SchemaName    string `json:"schemaName"`
	TypeName      string `json:"typeName"`
	FieldName     string `json:"fieldName"`
	FieldTypeName string `json:"fieldTypeName"`
}

// CacheMetrics is the response for Cache metrics command
// See https://apacheignite.readme.io/v1.9/docs/rest-api#section-cache-metrics for more details
type CacheMetrics struct {
	CreateTime *JavaTime `json:"createTime"`
	Hits       int64     `json:"hits"`
	Misses     int64     `json:"misses"`
	ReadTime   *JavaTime `json:"readTime"`
	Reads      int64     `json:"reads"`
	WriteTime  *JavaTime `json:"writeTime"`
	Writes     int64     `json:"writes"`
}
