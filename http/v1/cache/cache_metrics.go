package cache

import (
	"encoding/json"
	"net/url"

	"github.com/pkg/errors"

	core "github.com/amsokol/go-ignite-client/http"
	"github.com/amsokol/go-ignite-client/http/v1/exec"
)

// GetCacheMetrics shows metrics for Ignite cache
// See https://apacheignite.readme.io/v1.9/docs/rest-api#section-cache-metrics for more details
func (c *ManagementImpl) GetCacheMetrics(e exec.Executer, cache string, destID string) (metrics core.CacheMetrics, nodeID string, token string, err error) {
	v := url.Values{}
	v.Add("cmd", "cache")
	if len(cache) > 0 {
		v.Add("cacheName", cache)
	}
	if len(destID) > 0 {
		v.Add("destId", destID)
	}

	b, nodeID, token, err := e.Execute(v)
	if err != nil {
		return metrics, nodeID, token, err
	}

	err = json.Unmarshal(b, &metrics)
	if err != nil {
		return metrics, nodeID, token, errors.Wrap(err, "Can't unmarshal 'response' to CacheMetrics")
	}

	return metrics, nodeID, token, nil
}
