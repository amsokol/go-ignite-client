package v10

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/pkg/errors"

	"github.com/amsokol/go-ignite-client/http/types"
	"github.com/amsokol/go-ignite-client/http/v1/client"
)

// responseVersion is response for `log` command
// See https://apacheignite.readme.io/v1.0/docs/rest-api#log for more details
type responseLog struct {
	response
	Log string `json:"response"`
}

// Log command shows server logs
// See https://apacheignite.readme.io/v1.0/docs/rest-api#log for more details
func Log(c client.Client, path string, from int, to int) (string, types.SessionToken, error) {
	v := url.Values{}
	v.Add("cmd", "log")
	if len(path) > 0 {
		v.Add("path", path)
	}
	v.Add("from", strconv.FormatInt(int64(from), 10))
	v.Add("to", strconv.FormatInt(int64(to), 10))

	b, err := c.Execute(v)
	if err != nil {
		return "", types.SessionTokenNil, err
	}

	res := &responseLog{}
	err = json.Unmarshal(b, res)
	if err != nil {
		return "", types.SessionTokenNil, errors.Wrap(err, "Can't unmarshal respone to responseLog")
	}

	if c.IsFailed(res.SuccessStatus) {
		return "", types.SessionTokenNil, errors.New(c.GetError(res.SuccessStatus, res.Error))
	}

	return res.Log, res.SessionToken, nil
}
