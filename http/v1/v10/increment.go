package v10

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/pkg/errors"

	"github.com/amsokol/go-ignite-client/http/types"
	"github.com/amsokol/go-ignite-client/http/v1/client"
)

// responseIncrement is response for `incr` command
// See https://apacheignite.readme.io/v1.0/docs/rest-api#section-increment for more details
type responseIncrement struct {
	response
	AffinityNodeID string `json:"affinityNodeId"`
	Value          int64  `json:"response"`
}

// Increment command adds and gets current value of given atomic long
// See https://apacheignite.readme.io/v1.0/docs/rest-api#section-increment for more details
func Increment(c client.Client, cacheName string, key string, init *int64, delta int64) (int64, string, types.SessionToken, error) {
	v := url.Values{}
	v.Add("cmd", "incr")
	v.Add("cacheName", cacheName)
	v.Add("key", key)
	if init != nil {
		v.Add("init", strconv.FormatInt(int64(*init), 10))
	}
	v.Add("delta", strconv.FormatInt(int64(delta), 10))

	b, err := c.Execute(v)
	if err != nil {
		return 0, "", types.SessionTokenNil, err
	}

	res := &responseIncrement{}
	err = json.Unmarshal(b, res)
	if err != nil {
		return 0, "", types.SessionTokenNil, errors.Wrap(err, "Can't unmarshal respone to responseIncrement")
	}

	if c.IsFailed(res.SuccessStatus) {
		return 0, "", types.SessionTokenNil, errors.New(c.GetError(res.SuccessStatus, res.Error))
	}

	return res.Value, res.AffinityNodeID, res.SessionToken, nil
}
