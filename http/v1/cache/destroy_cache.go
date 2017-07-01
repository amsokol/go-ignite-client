package cache

import (
	"net/url"

	"github.com/amsokol/go-ignite-client/http/v1/exec"
)

// DestroyCache destroys cache with given name
// See https://apacheignite.readme.io/v1.9/docs/rest-api#section-destroy-cache for more details
func (c *Commands) DestroyCache(e exec.Executer, cache string) (token string, err error) {
	v := url.Values{}
	v.Add("cmd", "destcache")
	if len(cache) > 0 {
		v.Add("cacheName", cache)
	}

	_, _, token, err = e.Execute(v)
	if err != nil {
		return token, err
	}

	return token, nil
}
