package sql

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/pkg/errors"

	"github.com/amsokol/go-ignite-client/http/v1/exec"
)

// SQLQueryClose closes query resources
// See https://apacheignite.readme.io/v1.9/docs/rest-api#section-sql-query-close for more details
func (c *Commands) SQLQueryClose(e exec.Executer, queryID int64) (ok bool, token string, err error) {
	v := url.Values{}
	v.Add("cmd", "qrycls")
	v.Add("qryId", strconv.FormatInt(queryID, 10))

	b, _, token, err := e.Execute(v)
	if err != nil {
		return ok, token, err
	}

	err = json.Unmarshal(b, &ok)
	if err != nil {
		return ok, token, errors.Wrap(err, "Can't unmarshal respone to bool")
	}

	return ok, token, nil
}
