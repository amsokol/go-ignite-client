package sql

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/pkg/errors"

	core "github.com/amsokol/go-ignite-client/http"
	"github.com/amsokol/go-ignite-client/http/v1/exec"
)

// SQLFieldsQueryExecute runs sql fields query over cache.
// See https://apacheignite.readme.io/v1.9/docs/rest-api#section-sql-fields-query-execute for more details
func (c *Commands) SQLFieldsQueryExecute(e exec.Executer, cache string, pageSize int64, query string, args url.Values) (result core.SQLQueryResult, token string, err error) {
	if args == nil {
		args = url.Values{}
	}

	args.Add("cmd", "qryfldexe")
	if len(cache) > 0 {
		args.Add("cacheName", cache)
	}
	args.Add("qry", query)
	args.Add("pageSize", strconv.FormatInt(pageSize, 10))

	b, _, token, err := e.Execute(args)
	if err != nil {
		return result, token, err
	}

	err = json.Unmarshal(b, &result)
	if err != nil {
		return result, token, errors.Wrap(err, "Can't unmarshal respone to SQLQueryResult")
	}

	return result, token, nil
}
