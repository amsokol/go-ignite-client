package sql

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/pkg/errors"

	core "github.com/amsokol/go-ignite-client/http"
	"github.com/amsokol/go-ignite-client/http/v1/exec"
)

// SQLQueryFetch gets next page for the query
// See https://apacheignite.readme.io/v1.9/docs/rest-api#section-sql-query-fetch for more details
func (c *Commands) SQLQueryFetch(e exec.Executer, pageSize int64, queryID int64) (result core.SQLQueryResult, token string, err error) {
	v := url.Values{}
	v.Add("cmd", "qryfetch")
	v.Add("qryId", strconv.FormatInt(queryID, 10))
	v.Add("pageSize", strconv.FormatInt(pageSize, 10))

	b, _, token, err := e.Execute(v)
	if err != nil {
		return result, token, err
	}

	err = json.Unmarshal(b, &result)
	if err != nil {
		return result, token, errors.Wrap(err, "Can't unmarshal respone to SQLQueryResult")
	}

	return result, token, nil
}
