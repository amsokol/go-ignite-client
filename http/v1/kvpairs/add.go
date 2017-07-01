package kvpairs

import (
	"encoding/json"
	"net/url"

	"github.com/pkg/errors"

	"github.com/amsokol/go-ignite-client/http/v1/exec"
)

// Add command stores a given key-value pair in cache only if there isn't a previous mapping for it
// See https://apacheignite.readme.io/v1.9/docs/rest-api#section-add for for details
func (p *Commands) Add(e exec.Executer, cache string, key string, val string, destID string) (ok bool, nodeID string, token string, err error) {
	v := url.Values{}
	v.Add("cmd", "append")
	if len(cache) > 0 {
		v.Add("cacheName", cache)
	}
	v.Add("key", key)
	v.Add("val", val)
	if len(destID) > 0 {
		v.Add("destId", destID)
	}

	b, nodeID, token, err := e.Execute(v)
	if err != nil {
		return ok, nodeID, token, err
	}

	err = json.Unmarshal(b, &ok)
	if err != nil {
		return ok, nodeID, token, errors.Wrap(err, "Can't unmarshal respone to bool")
	}

	return ok, nodeID, token, nil
}
