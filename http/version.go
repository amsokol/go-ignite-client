package http

import (
	"encoding/json"
	"github.com/pkg/errors"
	"net/url"
)

// Version command shows current Ignite version.
//
// Request Parameters
// | name | type   | optional | description                  | example |
// |------|--------|----------|------------------------------|---------|
// | cmd  | string |          | Should be version lowercase. |         |
func (c *Client) Version() (Version, string, error) {
	v := url.Values{}
	v.Add("cmd", "version")

	b, err := c.execute(&v)
	if err != nil {
		return "", "", err
	}

	r := WrapperResponseVersion{}
	err = json.Unmarshal(b, &r)
	if err != nil {
		return "", "", errors.WithStack(errors.Wrap(err, "Can't unmarshal respone to WrapperResponseVersion"))
	}

	if r.SuccessStatus != successStatusSuccess {
		return "", "", errors.New(c.getError(r.SuccessStatus, r.Error))
	}

	return r.Version, r.SessionToken, nil
}
