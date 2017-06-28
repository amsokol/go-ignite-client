package v13

import (
	"encoding/json"
	"net/url"

	"github.com/pkg/errors"

	"github.com/amsokol/go-ignite-client/http/types"
	"github.com/amsokol/go-ignite-client/http/v1/client"
)

// responseSQLQueryClose is response for `qrycls` commands
// See https://apacheignite.readme.io/v1.3/docs/rest-api#section-sql-query-close for more details
type responseSQLQueryClose struct {
	SuccessStatus types.SuccessStatus `json:"successStatus"`
	Error         string              `json:"error"`
	Response      bool                `json:"response"`
	SessionToken  types.SessionToken  `json:"sessionToken"`
}

// SQLQueryClose closes query resources
// See https://apacheignite.readme.io/v1.3/docs/rest-api#section-sql-query-close for more details
func SQLQueryClose(c client.Client, queryID string) (bool, types.SessionToken, error) {
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

	if c.IsFailed(res.SuccessStatus) {
		return false, "", errors.New(c.GetError(res.SuccessStatus, res.Error))
	}

	return res.Response, res.SessionToken, nil
}
