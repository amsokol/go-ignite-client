package v10

import (
	"encoding/json"
	"net/url"

	"github.com/blang/semver"
	"github.com/pkg/errors"

	"github.com/amsokol/go-ignite-client/http/v1/internal"
)

// responseVersion is response for `version` command
// See https://apacheignite.readme.io/v1.0/docs/rest-api#section-version for more details
type responseVersion struct {
	SuccessStatus int64  `json:"successStatus"`
	Error         string `json:"error"`
	Version       string `json:"response"`
	SessionToken  string `json:"sessionToken"`
}

func (r *responseVersion) GetSuccessStatus() internal.SuccessStatus {
	return internal.SuccessStatus(r.SuccessStatus)
}

func (r *responseVersion) GetError() string {
	return r.Error
}

func (r *responseVersion) GetVersion() string {
	return r.Version
}

func (r *responseVersion) GetSessionToken() internal.SessionToken {
	return internal.SessionToken(r.SessionToken)
}

// Version command shows current Ignite version.
// See https://apacheignite.readme.io/v1.0/docs/rest-api#section-version for more details
func Version(c internal.Client) (internal.Version, internal.SessionToken, error) {
	v := url.Values{}
	v.Add("cmd", "version")

	b, err := c.Execute(v)
	if err != nil {
		return internal.Version{}, internal.SessionTokenNil, err
	}

	res := &responseVersion{}
	err = json.Unmarshal(b, res)
	if err != nil {
		return internal.Version{}, internal.SessionTokenNil, errors.Wrap(err, "Can't unmarshal respone to ResponseVersion")
	}

	if c.IsFailed(res.GetSuccessStatus()) {
		return internal.Version{}, internal.SessionTokenNil, errors.New(c.GetError(res.GetSuccessStatus(), res.GetError()))
	}

	sv, err := semver.Make(res.Version)
	if err != nil {
		return internal.Version{}, internal.SessionTokenNil, errors.Wrap(err, "Server returned version in unsupported format")
	}

	return internal.Version(sv), res.GetSessionToken(), nil
}
