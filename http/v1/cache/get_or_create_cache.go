package cache

import (
	"net/url"

	"github.com/amsokol/go-ignite-client/http/v1/exec"
)

// GetOrCreateCache creates cache with given name if it does not exist
// See https://apacheignite.readme.io/v1.9/docs/rest-api#section-get-or-create-cache for more details
func (c *Commands) GetOrCreateCache(e exec.Executer, cache string) (token string, err error) {
	v := url.Values{}
	v.Add("cmd", "getorcreate")
	if len(cache) > 0 {
		v.Add("cacheName", cache)
	}

	_, _, token, err = e.Execute(v)
	if err != nil {
		return token, err
	}

	return token, nil
}
