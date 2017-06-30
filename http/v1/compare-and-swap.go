package v1

import (
	"encoding/json"
	"net/url"

	"github.com/pkg/errors"
)

// CompareAndSwap stores given key-value pair in cache only if the previous value is equal to the expected value passed in
// See https://apacheignite.readme.io/v1.9/docs/rest-api#section-compare-and-swap for details
func (c *client) CompareAndSwap(cache string, key string, val string, val2 string, destID string) (ok bool, nodeID string, token string, err error) {
	v := url.Values{}
	v.Add("cmd", "cas")
	if len(cache) > 0 {
		v.Add("cacheName", cache)
	}
	v.Add("key", key)
	v.Add("val", val)
	v.Add("val2", val2)
	if len(destID) > 0 {
		v.Add("destId", destID)
	}

	b, nodeID, token, err := c.execute(v)
	if err != nil {
		return ok, nodeID, token, err
	}

	err = json.Unmarshal(b, &ok)
	if err != nil {
		return ok, nodeID, token, errors.Wrap(err, "Can't unmarshal respone to bool")
	}

	return ok, nodeID, token, nil
}
