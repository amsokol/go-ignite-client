package v13

import (
	"github.com/amsokol/go-ignite-client/http/v1/internal"
	"github.com/amsokol/go-ignite-client/http/v1/v10"
)

// Version command shows current Ignite version.
// See https://apacheignite.readme.io/v1.3/docs/rest-api#section-version for more details
func Version(c internal.Client) (internal.Version, internal.SessionToken, error) {
	return v10.Version(c)
}
