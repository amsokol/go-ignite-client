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
	SuccessStatus int64  `json:"successStatus"`
	Error         string `json:"error"`
	Version       string `json:"response"`
	SessionToken  string `json:"sessionToken"`
}

func (r *responseVersion) GetSuccessStatus() common.SuccessStatus {
	return common.SuccessStatus(r.SuccessStatus)
}

func (r *responseVersion) GetError() string {
	return r.Error
}

func (r *responseVersion) GetVersion() string {
	return r.Version
}

func (r *responseVersion) GetSessionToken() common.SessionToken {
	return common.SessionToken(r.SessionToken)
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

	if c.IsFailed(res.GetSuccessStatus()) {
		return common.Version{}, common.SessionTokenNil, errors.New(c.GetError(res.GetSuccessStatus(), res.GetError()))
	}

	sv, err := semver.Make(res.Version)
	if err != nil {
		return common.Version{}, common.SessionTokenNil, errors.Wrap(err, "Server returned version in unsupported format")
	}

	return common.Version(sv), res.GetSessionToken(), nil
}
