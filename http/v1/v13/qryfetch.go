package v13

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/pkg/errors"

	"github.com/amsokol/go-ignite-client/http/v1/common"
)

// responseSQLQueryFetch is response for `qryfetch`, commands
// See https://apacheignite.readme.io/v1.3/docs/rest-api#section-sql-query-fetch for more details
type responseSQLQueryFetch struct {
	SuccessStatus int64          `json:"successStatus"`
	Error         string         `json:"error"`
	Response      sqlQueryResult `json:"response"`
	SessionToken  string         `json:"sessionToken"`
}

// GetSuccessStatus implements common.ResponseSQLQueryFetch interface
func (r *responseSQLQueryFetch) GetSuccessStatus() common.SuccessStatus {
	return common.SuccessStatus(r.SuccessStatus)
}

// GetError implements common.ResponseSQLQueryFetch interface
func (r *responseSQLQueryFetch) GetError() string {
	return r.Error
}

// GetSessionToken implements common.ResponseSQLQueryFetch interface
func (r *responseSQLQueryFetch) GetSessionToken() common.SessionToken {
	return common.SessionToken(r.SessionToken)
}

// GetResponse implements common.ResponseSQLQueryFetch interface
func (r *responseSQLQueryFetch) GetResponse() common.SQLQueryResult {
	return &r.Response
}

// sqlQueryResult is body of response for `qryfetch`, command
// See https://apacheignite.readme.io/v1.3/docs/rest-api#section-sql-query-fetch for more details
type sqlQueryResult struct {
	Items          [][]interface{} `json:"items"`
	Last           bool            `json:"last"`
	QueryID        int64           `json:"queryId"`
	FieldsMetadata []fieldMetadata `json:"fieldsMetadata"`
}

// GetItems implements common.SQLQueryResult interface
func (r *sqlQueryResult) GetItems() [][]interface{} {
	return r.Items
}

// GetLast implements common.SQLQueryResult interface
func (r *sqlQueryResult) GetLast() bool {
	return r.Last
}

// GetQueryID implements common.SQLQueryResult interface
func (r *sqlQueryResult) GetQueryID() int64 {
	return r.QueryID
}

// GetFieldsMetadata implements common.SQLQueryResult interface
func (r *sqlQueryResult) GetFieldsMetadata() []common.FieldMetadata {
	size := len(r.FieldsMetadata)
	res := make([]common.FieldMetadata, size, size)
	for i, v := range r.FieldsMetadata {
		res[i] = &v
	}
	return res
}

// fieldMetadata is column list
// See https://apacheignite.readme.io/v1.3/docs/rest-api#section-sql-fields-query-execute for more details
type fieldMetadata struct {
	SchemaName    string `json:"schemaName"`
	TypeName      string `json:"typeName"`
	FieldName     string `json:"fieldName"`
	FieldTypeName string `json:"fieldTypeName"`
}

// GetSchemaName implements common.FieldMetadata interface
func (m *fieldMetadata) GetSchemaName() string {
	return m.SchemaName
}

// GetTypeName implements common.FieldMetadata interface
func (m *fieldMetadata) GetTypeName() string {
	return m.TypeName
}

// GetFieldName implements common.FieldMetadata interface
func (m *fieldMetadata) GetFieldName() string {
	return m.FieldName
}

// GetFieldTypeName implements common.FieldMetadata interface
func (m *fieldMetadata) GetFieldTypeName() string {
	return m.FieldTypeName
}

// SQLQueryFetch gets next page for the query
// See https://apacheignite.readme.io/v1.3/docs/rest-api#section-sql-query-fetch for more details
func SQLQueryFetch(c common.Client, pageSize int64, queryID string) (common.SQLQueryResult, common.SessionToken, error) {
	v := url.Values{}
	v.Add("cmd", "qryfetch")
	v.Add("qryId", queryID)
	v.Add("pageSize", strconv.FormatInt(pageSize, 10))

	b, err := c.Execute(v)
	if err != nil {
		return nil, "", err
	}

	res := responseSQLQueryFetch{}
	err = json.Unmarshal(b, &res)
	if err != nil {
		return nil, "", errors.Wrap(err, "Can't unmarshal respone to WrapperResponse")
	}

	if c.IsFailed(res.GetSuccessStatus()) {
		return nil, "", errors.New(c.GetError(res.GetSuccessStatus(), res.GetError()))
	}

	return res.GetResponse(), res.GetSessionToken(), nil
}
