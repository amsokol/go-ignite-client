package v13

import (
	"encoding/json"
	"github.com/pkg/errors"
	"net/url"
	"strconv"

	"github.com/amsokol/go-ignite-client/http/v1/internal"
)

// responseSQLFieldsQueryExecute is response for `qryfldexe`, commands
// See https://apacheignite.readme.io/v1.3/docs/rest-api#section-sql-fields-query-execute for more details
type responseSQLFieldsQueryExecute struct {
	SuccessStatus int64          `json:"successStatus"`
	Error         string         `json:"error"`
	Response      sqlQueryResult `json:"response"`
	SessionToken  string         `json:"sessionToken"`
}

// GetSuccessStatus implements internal.ResponseSQLFieldsQueryExecute interface
func (r *responseSQLFieldsQueryExecute) GetSuccessStatus() internal.SuccessStatus {
	return internal.SuccessStatus(r.SuccessStatus)
}

// GetError implements internal.ResponseSQLFieldsQueryExecute interface
func (r *responseSQLFieldsQueryExecute) GetError() string {
	return r.Error
}

// GetSessionToken implements internal.ResponseSQLFieldsQueryExecute interface
func (r *responseSQLFieldsQueryExecute) GetSessionToken() internal.SessionToken {
	return internal.SessionToken(r.SessionToken)
}

// Response implements internal.ResponseSQLFieldsQueryExecute interface
func (r *responseSQLFieldsQueryExecute) GetResponse() internal.SQLQueryResult {
	return &r.Response
}

// SQLFieldsQueryExecute runs sql fields query over cache.
// See https://apacheignite.readme.io/v1.3/docs/rest-api#section-sql-fields-query-execute for more details
func SQLFieldsQueryExecute(c internal.Client, cacheName string, pageSize int64, query string, args url.Values) (internal.SQLQueryResult, internal.SessionToken, error) {
	args.Add("cmd", "qryfldexe")
	args.Add("cacheName", cacheName)
	args.Add("qry", query)
	args.Add("pageSize", strconv.FormatInt(pageSize, 10))

	b, err := c.Execute(args)
	if err != nil {
		return nil, "", err
	}

	res := responseSQLFieldsQueryExecute{}
	err = json.Unmarshal(b, &res)
	if err != nil {
		return nil, "", errors.Wrap(err, "Can't unmarshal respone to WrapperResponse")
	}

	if c.IsFailed(res.GetSuccessStatus()) {
		return nil, "", errors.New(c.GetError(res.GetSuccessStatus(), res.GetError()))
	}

	return res.GetResponse(), res.GetSessionToken(), nil
}
