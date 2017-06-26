package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/blang/semver"
	"github.com/pkg/errors"
)

// ResponseVersion is response for `version` command
// See https://apacheignite.readme.io/v1.0/docs/rest-api#section-version for more details
type ResponseVersion struct {
	SuccessStatus int    `json:"successStatus"`
	Error         string `json:"error"`
	Version       string `json:"response"`
}

// Version command shows current Ignite version.
// See https://apacheignite.readme.io/v1.0/docs/rest-api#section-version for more details
func Version(servers []string, username string, password string) (*semver.Version, error) {
	v := url.Values{}
	v.Add("cmd", "version")

	// TODO: add round-robin to select node
	req, err := http.NewRequest("POST", servers[0], strings.NewReader(v.Encode()))
	if err != nil {
		return nil, errors.Wrap(err, "Can't create new POST http.Request")
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	if len(username) > 0 {
		req.SetBasicAuth(username, password)
	}

	r, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "Can't Do HTTP request by DefaultClient")
	}

	b, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		return nil, errors.Wrap(err, "Can't read bytes from HTTP response body")
	}

	res := ResponseVersion{}

	err = json.Unmarshal(b, &res)
	if err != nil {
		return nil, errors.Wrap(err, "Can't unmarshal respone to ResponseVersion")
	}

	if res.SuccessStatus != 0 {
		return nil, errors.New(res.Error)
	}

	ver, err := semver.Make(res.Version)
	if err != nil {
		return nil, errors.Wrap(err, "Server returns version in unsupported format")
	}

	return &ver, err
}
