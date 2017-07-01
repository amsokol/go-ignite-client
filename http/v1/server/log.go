package server

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/pkg/errors"

	"github.com/amsokol/go-ignite-client/http/v1/exec"
)

// GetLog command shows server logs
// See https://apacheignite.readme.io/v1.9/docs/rest-api#log for more details
func (c *Commands) GetLog(e exec.Executer, path string, from *int, to *int) (log string, token string, err error) {
	v := url.Values{}
	v.Add("cmd", "log")
	if len(path) > 0 {
		v.Add("path", path)
	}
	if from != nil {
		v.Add("from", strconv.FormatInt(int64(*from), 10))
	}
	if from != nil {
		v.Add("to", strconv.FormatInt(int64(*to), 10))
	}

	b, _, token, err := e.Execute(v)
	if err != nil {
		return log, token, err
	}

	err = json.Unmarshal(b, &log)
	if err != nil {
		return log, token, errors.Wrap(err, "Can't unmarshal respone to string")
	}

	return log, token, nil
}
