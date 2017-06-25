package http

import (
	"encoding/json"
	"net/url"

	"github.com/pkg/errors"
)

// QryCls closes query resources
func (c *Client) QryCls(queryId string) (bool, string, error) {
	v := url.Values{}
	v.Add("cmd", "qrycls")
	v.Add("qryId", queryId)

	b, err := c.execute(&v)
	if err != nil {
		return false, "", err
	}

	r := WrapperResponseBinary{}
	err = json.Unmarshal(b, &r)
	if err != nil {
		return false, "", errors.WithStack(errors.Wrap(err, "Can't unmarshal respone to WrapperBinaryResponse"))
	}

	if r.SuccessStatus != successStatusSuccess {
		return false, "", errors.WithStack(errors.New(c.getError(r.SuccessStatus, r.Error)))
	}

	return r.Response, r.SessionToken, nil
}
