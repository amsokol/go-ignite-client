package http

import (
	"encoding/json"
	"net/url"

	"github.com/pkg/errors"
)

// QryFetch gets next page for the query
// See https://apacheignite.readme.io/docs/rest-api#section-sql-query-fetch for more details
func (c *Client) QryFetch(queryID string) (*Response, string, error) {
	v := url.Values{}
	v.Add("cmd", "qryfetch")
	v.Add("qryId", queryID)
	v.Add("pageSize", c.ConnectionInfo.PageSizeStr)

	b, err := c.execute(&v)
	if err != nil {
		return nil, "", err
	}

	r := WrapperResponse{}
	err = json.Unmarshal(b, &r)
	if err != nil {
		return nil, "", errors.WithStack(errors.Wrap(err, "Can't unmarshal respone to WrapperResponse"))
	}

	if r.SuccessStatus != successStatusSuccess {
		return nil, "", errors.WithStack(errors.New(c.getError(r.SuccessStatus, r.Error)))
	}

	return &r.Response, r.SessionToken, nil
}
