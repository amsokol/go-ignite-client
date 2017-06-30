package v1

import (
	"encoding/json"
	"net/url"

	"github.com/pkg/errors"
)

// Replace stores a given key-value pair in cache only if there is a previous mapping for it
// See https://apacheignite.readme.io/v1.9/docs/rest-api#section-replace for more details
func (c *client) Replace(cache string, key string, val string, destID string) (ok bool, nodeID string, token string, err error) {
	v := url.Values{}
	v.Add("cmd", "rep")
	if len(cache) > 0 {
		v.Add("cacheName", cache)
	}
	v.Add("key", key)
	v.Add("val", val)
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
