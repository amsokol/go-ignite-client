package http

import (
	"encoding/json"
	"github.com/pkg/errors"
	"net/url"
	"strconv"
)

// QryFldExe runs sql fields query over cache.
// See https://apacheignite.readme.io/docs/rest-api#section-sql-fields-query-execute for more details
func (c *Client) QryFldExe(query string, v url.Values) (*Response, string, error) {
	v.Add("cmd", "qryfldexe")
	v.Add("cacheName", c.ConnectionInfo.Cache)
	v.Add("qry", query)
	v.Add("pageSize", strconv.FormatInt(c.ConnectionInfo.PageSize, 10))

	b, err := c.execute(v)
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
