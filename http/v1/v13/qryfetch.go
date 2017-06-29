package v13

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/pkg/errors"

	"github.com/amsokol/go-ignite-client/http/types"
	"github.com/amsokol/go-ignite-client/http/v1/client"
)

// responseSQLQueryFetch is response for `qryfetch`, commands
// See https://apacheignite.readme.io/v1.3/docs/rest-api#section-sql-query-fetch for more details
type responseSQLQueryFetch struct {
	response
	Response types.SQLQueryResult `json:"response"`
}

// SQLQueryFetch gets next page for the query
// See https://apacheignite.readme.io/v1.3/docs/rest-api#section-sql-query-fetch for more details
func SQLQueryFetch(c client.Client, pageSize int64, queryID string) (*types.SQLQueryResult, types.SessionToken, error) {
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

	if c.IsFailed(res.SuccessStatus) {
		return nil, "", errors.New(c.GetError(res.SuccessStatus, res.Error))
	}

	return &res.Response, res.SessionToken, nil
}
