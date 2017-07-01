package kvpairs

import (
	"encoding/json"
	"net/url"

	"github.com/pkg/errors"

	"github.com/amsokol/go-ignite-client/http/v1/exec"
)

// Append appends a line for value which is associated with key
// See https://apacheignite.readme.io/v1.9/docs/rest-api#section-append for more details
func (p *Commands) Append(e exec.Executer, cache string, key string, val string, destID string) (ok bool, nodeID string, token string, err error) {
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
