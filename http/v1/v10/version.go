package v10

import (
	"encoding/json"
	"net/url"

	"github.com/blang/semver"
	"github.com/pkg/errors"

	"github.com/amsokol/go-ignite-client/http/v1/common"
)

// responseVersion is response for `version` command
// See https://apacheignite.readme.io/v1.0/docs/rest-api#section-version for more details
type responseVersion struct {
	SuccessStatus common.SuccessStatus `json:"successStatus"`
	Error         string               `json:"error"`
	Version       string               `json:"response"`
	SessionToken  common.SessionToken  `json:"sessionToken"`
}

// Version command shows current Ignite version.
// See https://apacheignite.readme.io/v1.0/docs/rest-api#section-version for more details
func Version(c common.Client) (common.Version, common.SessionToken, error) {
	v := url.Values{}
	v.Add("cmd", "version")

	b, err := c.Execute(v)
	if err != nil {
		return common.Version{}, common.SessionTokenNil, err
	}

	res := &responseVersion{}
	err = json.Unmarshal(b, res)
	if err != nil {
		return common.Version{}, common.SessionTokenNil, errors.Wrap(err, "Can't unmarshal respone to ResponseVersion")
	}

	if c.IsFailed(res.SuccessStatus) {
		return common.Version{}, common.SessionTokenNil, errors.New(c.GetError(res.SuccessStatus, res.Error))
	}

	sv, err := semver.Make(res.Version)
	if err != nil {
		return common.Version{}, common.SessionTokenNil, errors.Wrap(err, "Server returned version in unsupported format")
	}

	return common.Version(sv), res.SessionToken, nil
}
