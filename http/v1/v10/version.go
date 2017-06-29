package v10

import (
	"encoding/json"
	"net/url"

	"github.com/blang/semver"
	"github.com/pkg/errors"

	"github.com/amsokol/go-ignite-client/http/types"
	"github.com/amsokol/go-ignite-client/http/v1/client"
)

// responseVersion is response for `version` command
// See https://apacheignite.readme.io/v1.0/docs/rest-api#section-version for more details
type responseVersion struct {
	response
	Version string `json:"response"`
}

// Version command shows current Ignite version.
// See https://apacheignite.readme.io/v1.0/docs/rest-api#section-version for more details
func Version(c client.Client) (types.Version, types.SessionToken, error) {
	v := url.Values{}
	v.Add("cmd", "version")

	b, err := c.Execute(v)
	if err != nil {
		return types.Version{}, types.SessionTokenNil, err
	}

	res := &responseVersion{}
	err = json.Unmarshal(b, res)
	if err != nil {
		return types.Version{}, types.SessionTokenNil, errors.Wrap(err, "Can't unmarshal respone to responseVersion")
	}

	if c.IsFailed(res.SuccessStatus) {
		return types.Version{}, types.SessionTokenNil, errors.New(c.GetError(res.SuccessStatus, res.Error))
	}

	sv, err := semver.Make(res.Version)
	if err != nil {
		return types.Version{}, types.SessionTokenNil, errors.Wrap(err, "Server returned version in unsupported format")
	}

	return types.Version(sv), res.SessionToken, nil
}
