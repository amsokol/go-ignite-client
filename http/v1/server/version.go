package server

import (
	"encoding/json"
	"net/url"

	"github.com/pkg/errors"

	"github.com/amsokol/go-ignite-client/http/v1/exec"
)

// GetVersion command shows current Ignite version.
// See https://apacheignite.readme.io/v1.9/docs/rest-api#section-version for more details
func (u *Commands) GetVersion(e exec.Executer) (version string, token string, err error) {
	v := url.Values{}
	v.Add("cmd", "version")

	b, _, token, err := e.Execute(v)
	if err != nil {
		return version, token, err
	}

	err = json.Unmarshal(b, &version)
	if err != nil {
		return version, token, errors.Wrap(err, "Can't unmarshal respone to string")
	}

	return version, token, nil
}
