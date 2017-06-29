package v10

import (
	"encoding/json"
	"net/url"

	"github.com/pkg/errors"

	"github.com/amsokol/go-ignite-client/http/types"
	"github.com/amsokol/go-ignite-client/http/v1/client"
)

// responseCacheMetrics is response for `cache` command
// See https://apacheignite.readme.io/v1.0/docs/rest-api#section-cache-metrics for more details
type responseCacheMetrics struct {
	response
	AffinityNodeID string             `json:"affinityNodeId"`
	CacheMetrics   types.CacheMetrics `json:"response"`
}

// CacheMetrics shows metrics for Ignite cache
// See https://apacheignite.readme.io/v1.0/docs/rest-api#section-cache-metrics for more details
func CacheMetrics(c client.Client, cacheName string, destID string) (
	metrics types.CacheMetrics, affinityNodeID string, sessionToken types.SessionToken, err error) {
	v := url.Values{}
	v.Add("cmd", "cache")
	if len(cacheName) > 0 {
		v.Add("cacheName", cacheName)
	}
	if len(destID) > 0 {
		v.Add("destId", destID)
	}

	res := responseCacheMetrics{}

	b, err := c.Execute(v)
	if err != nil {
		return res.CacheMetrics, "", types.SessionTokenNil, err
	}

	err = json.Unmarshal(b, &res)
	if err != nil {
		return res.CacheMetrics, "", types.SessionTokenNil, errors.Wrap(err, "Can't unmarshal respone to responseDecrement")
	}

	if c.IsFailed(res.SuccessStatus) {
		return res.CacheMetrics, "", types.SessionTokenNil, errors.New(c.GetError(res.SuccessStatus, res.Error))
	}

	return res.CacheMetrics, res.AffinityNodeID, res.SessionToken, nil
}
