package v1

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/pkg/errors"
)

// SQLQueryClose closes query resources
// See https://apacheignite.readme.io/v1.3/docs/rest-api#section-sql-query-close for more details
func (c *client) SQLQueryClose(queryID int64) (ok bool, token string, err error) {
	v := url.Values{}
	v.Add("cmd", "qrycls")
	v.Add("qryId", strconv.FormatInt(queryID, 10))

	b, _, token, err := c.execute(v)
	if err != nil {
		return ok, token, err
	}

	err = json.Unmarshal(b, &ok)
	if err != nil {
		return ok, token, errors.Wrap(err, "Can't unmarshal respone to bool")
	}

	return ok, token, nil
}
