package http

const (
	successStatusSuccess             = 0
	successStatusFailed              = 1
	successStatusAuthorizationFailed = 2
	successStatusSecurityCheckFailed = 3
	successStatusUnknown             = 4
)

var successStatusMsg = []string{"success", "failed", "authorization failed", "security check failed", "unknown status"}

// ConnectionInfo contains Ignite cluster connection information
type ConnectionInfo struct {
	Servers     []string `json:"servers"`
	Username    string   `json:"username"`
	Password    string   `json:"password"`
	Cache       string   `json:"cache"`
	PageSize    int64    `json:"pageSize"`
	PageSizeStr string
}

type WrapperResponseVersion struct {
	SuccessStatus int     `json:"successStatus"`
	Error         string  `json:"error"`
	Version       Version `json:"response"`
	SessionToken  string  `json:"sessionToken"`
}

type Version string

type WrapperResponse struct {
	SuccessStatus int      `json:"successStatus"`
	Error         string   `json:"error"`
	Response      Response `json:"response"`
	SessionToken  string   `json:"sessionToken"`
}

type Response struct {
	Items          [][]interface{} `json:"items"`
	Last           bool            `json:"last"`
	QueryID        int64           `json:"queryId"`
	FieldsMetadata []FieldMetadata `json:"fieldsMetadata"`
}

type FieldMetadata struct {
	SchemaName    string `json:"schemaName"`
	TypeName      string `json:"typeName"`
	FieldName     string `json:"fieldName"`
	FieldTypeName string `json:"fieldTypeName"`
}

type WrapperResponseBinary struct {
	SuccessStatus int    `json:"successStatus"`
	Error         string `json:"error"`
	Response      bool   `json:"response"`
	SessionToken  string `json:"sessionToken"`
}
