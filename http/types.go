package http

// See https://apacheignite.readme.io/docs/rest-api#section-returned-value for more details
const (
	successStatusSuccess             = 0
	successStatusFailed              = 1
	successStatusAuthorizationFailed = 2
	successStatusSecurityCheckFailed = 3
	successStatusUnknown             = 4
)

// See https://apacheignite.readme.io/docs/rest-api#section-returned-value for more details
var successStatusMsg = []string{"success", "failed", "authorization failed", "security check failed", "unknown status"}

// ConnectionInfo contains Ignite cluster connection information user should provide when open connection
type ConnectionInfo struct {
	Servers     []string `json:"servers"`
	Username    string   `json:"username"`
	Password    string   `json:"password"`
	Cache       string   `json:"cache"`
	PageSize    int64    `json:"pageSize"`
	PageSizeStr string
}

// WrapperResponseVersion is response for `version` command
// See https://apacheignite.readme.io/docs/rest-api#section-version for more details
type WrapperResponseVersion struct {
	SuccessStatus int     `json:"successStatus"`
	Error         string  `json:"error"`
	Version       Version `json:"response"`
	SessionToken  string  `json:"sessionToken"`
}

// Version is alias for `string`
type Version string

// WrapperResponse is response for `qryfldexe`, `qryfetch` commands
// See https://apacheignite.readme.io/docs/rest-api#section-sql-fields-query-execute for more details
// See https://apacheignite.readme.io/docs/rest-api#section-sql-query-fetch for more details
type WrapperResponse struct {
	SuccessStatus int      `json:"successStatus"`
	Error         string   `json:"error"`
	Response      Response `json:"response"`
	SessionToken  string   `json:"sessionToken"`
}

// Response is body of response for `qryfldexe`, `qryfetch` commands
// See https://apacheignite.readme.io/docs/rest-api#section-sql-fields-query-execute for more details
// See https://apacheignite.readme.io/docs/rest-api#section-sql-query-fetch for more details
type Response struct {
	Items          [][]interface{} `json:"items"`
	Last           bool            `json:"last"`
	QueryID        int64           `json:"queryId"`
	FieldsMetadata []FieldMetadata `json:"fieldsMetadata"`
}

// FieldMetadata is column list
// See https://apacheignite.readme.io/docs/rest-api#section-sql-fields-query-execute for more details
// See https://apacheignite.readme.io/docs/rest-api#section-sql-query-fetch for more details
type FieldMetadata struct {
	SchemaName    string `json:"schemaName"`
	TypeName      string `json:"typeName"`
	FieldName     string `json:"fieldName"`
	FieldTypeName string `json:"fieldTypeName"`
}

// WrapperResponseBinary is response for `qrycls` commands
// See https://apacheignite.readme.io/docs/rest-api#section-sql-query-close for more details
type WrapperResponseBinary struct {
	SuccessStatus int    `json:"successStatus"`
	Error         string `json:"error"`
	Response      bool   `json:"response"`
	SessionToken  string `json:"sessionToken"`
}
