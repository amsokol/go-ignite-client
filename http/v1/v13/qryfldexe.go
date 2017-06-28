package v13

import (
	"encoding/json"
	"github.com/pkg/errors"
	"net/url"
	"strconv"

	"github.com/amsokol/go-ignite-client/http/types"
	"github.com/amsokol/go-ignite-client/http/v1/client"
)

// responseSQLFieldsQueryExecute is response for `qryfldexe`, commands
// See https://apacheignite.readme.io/v1.3/docs/rest-api#section-sql-fields-query-execute for more details
type responseSQLFieldsQueryExecute struct {
	SuccessStatus types.SuccessStatus  `json:"successStatus"`
	Error         string               `json:"error"`
	Response      types.SQLQueryResult `json:"response"`
	SessionToken  types.SessionToken   `json:"sessionToken"`
}

// SQLFieldsQueryExecute runs sql fields query over cache.
// See https://apacheignite.readme.io/v1.3/docs/rest-api#section-sql-fields-query-execute for more details
func SQLFieldsQueryExecute(c client.Client, cacheName string, pageSize int64, query string, args url.Values) (*types.SQLQueryResult, types.SessionToken, error) {
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

	if c.IsFailed(res.SuccessStatus) {
		return nil, "", errors.New(c.GetError(res.SuccessStatus, res.Error))
	}

	return &res.Response, res.SessionToken, nil
}
