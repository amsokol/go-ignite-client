package v1

import (
	"encoding/json"
	"net/url"

	"github.com/pkg/errors"
)

// Version command shows current Ignite version.
// See https://apacheignite.readme.io/v1.0/docs/rest-api#section-version for more details
func (c *client) GetVersion() (version string, token string, err error) {
	v := url.Values{}
	v.Add("cmd", "version")

	b, _, token, err := c.execute(v)
	if err != nil {
		return version, token, err
	}

	err = json.Unmarshal(b, &version)
	if err != nil {
		return version, token, errors.Wrap(err, "Can't unmarshal respone to string")
	}

	return version, token, nil
}
