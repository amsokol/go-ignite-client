package v1

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/pkg/errors"
)

// Increment command adds and gets current value of given atomic long
// See https://apacheignite.readme.io/v1.0/docs/rest-api#section-increment for more details
func (c *client) Increment(cache string, key string, init *int64, delta int64) (value int64, nodeID string, token string, err error) {
	v := url.Values{}
	v.Add("cmd", "incr")
	v.Add("cacheName", cache)
	v.Add("key", key)
	if init != nil {
		v.Add("init", strconv.FormatInt(int64(*init), 10))
	}
	v.Add("delta", strconv.FormatInt(int64(delta), 10))

	b, nodeID, token, err := c.execute(v)
	if err != nil {
		return value, nodeID, token, err
	}

	err = json.Unmarshal(b, &value)
	if err != nil {
		return value, nodeID, token, errors.Wrap(err, "Can't unmarshal respone to responseIncrement")
	}

	return value, nodeID, token, nil
}
