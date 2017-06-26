package v13

import (
	"encoding/json"
	"net/url"

	"github.com/pkg/errors"

	"github.com/amsokol/go-ignite-client/http/v1/common"
)

// requestSQLQueryClose is request for `qrycls` commands
// See https://apacheignite.readme.io/v1.3/docs/rest-api#section-sql-query-close for more details
type requestSQLQueryClose struct {
	queryID string
}

// GetQueryID implements common.RequestSQLQueryClose interface
func (r *requestSQLQueryClose) GetQueryID() string {
	return r.queryID
}

// responseSQLQueryClose is response for `qrycls` commands
// See https://apacheignite.readme.io/v1.3/docs/rest-api#section-sql-query-close for more details
type responseSQLQueryClose struct {
	SuccessStatus int64  `json:"successStatus"`
	Error         string `json:"error"`
	Response      bool   `json:"response"`
	SessionToken  string `json:"sessionToken"`
}

// GetSuccessStatus implements common.ResponseSQLQueryClose interface
func (r *responseSQLQueryClose) GetSuccessStatus() common.SuccessStatus {
	return common.SuccessStatus(r.SuccessStatus)
}

// GetError implements common.ResponseSQLQueryClose interface
func (r *responseSQLQueryClose) GetError() string {
	return r.Error
}

// GetSessionToken implements common.ResponseSQLQueryClose interface
func (r *responseSQLQueryClose) GetSessionToken() common.SessionToken {
	return common.SessionToken(r.SessionToken)
}

// GetResponse implements common.ResponseSQLQueryClose interface
func (r *responseSQLQueryClose) GetResponse() bool {
	return r.Response
}

// SQLQueryClose closes query resources
// See https://apacheignite.readme.io/v1.3/docs/rest-api#section-sql-query-close for more details
func SQLQueryClose(c common.Client, queryID string) (bool, common.SessionToken, error) {
	v := url.Values{}
	v.Add("cmd", "qrycls")
	v.Add("qryId", queryID)

	b, err := c.Execute(v)
	if err != nil {
		return false, "", err
	}

	res := responseSQLQueryClose{}
	err = json.Unmarshal(b, &res)
	if err != nil {
		return false, "", errors.Wrap(err, "Can't unmarshal respone to ResponseSqlQueryClose")
	}

	if c.IsFailed(res.GetSuccessStatus()) {
		return false, "", errors.New(c.GetError(res.GetSuccessStatus(), res.GetError()))
	}

	return res.GetResponse(), res.GetSessionToken(), nil
}
