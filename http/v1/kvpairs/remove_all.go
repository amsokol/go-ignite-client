package kvpairs

import (
	"encoding/json"
	"net/url"

	"github.com/pkg/errors"

	"github.com/amsokol/go-ignite-client/http/v1/exec"
)

// RemoveAll removes given key mappings from cache
// See https://apacheignite.readme.io/v1.9/docs/rest-api#section-remove-all for details
func (c *Commands) RemoveAll(e exec.Executer, cache string, args url.Values, destID string) (ok bool, nodeID string, token string, err error) {
	if args == nil {
		args = url.Values{}
	}

	args.Add("cmd", "rmvall")
	if len(cache) > 0 {
		args.Add("cacheName", cache)
	}
	if len(destID) > 0 {
		args.Add("destId", destID)
	}

	b, nodeID, token, err := e.Execute(args)
	if err != nil {
		return ok, nodeID, token, err
	}

	err = json.Unmarshal(b, &ok)
	if err != nil {
		return ok, nodeID, token, errors.Wrap(err, "Can't unmarshal respone to bool")
	}

	return ok, nodeID, token, nil
}
