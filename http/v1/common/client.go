package common

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

// See https://apacheignite.readme.io/docs/rest-api#section-returned-value for more details
const (
	successStatusSuccess             = 0
	successStatusFailed              = 1
	successStatusAuthorizationFailed = 2
	successStatusSecurityCheckFailed = 3
	successStatusUnknown             = 4
)

// See https://apacheignite.readme.io/docs/rest-api#section-returned-value for more details
var successStatusMsg = []string{"success", "failed", "authorization failed", "security check failed", "unknown status"}

type client struct {
	servers  []string
	username string
	password string
}

// Open returns client
func Open(servers []string, username string, password string) Client {
	return &client{servers: servers, username: username, password: password}
}

// Execute implements http.CommandExecutor
func (c *client) Execute(v url.Values) ([]byte, error) {
	// TODO: add round-robin to select node
	req, err := http.NewRequest("POST", c.servers[0], strings.NewReader(v.Encode()))
	if err != nil {
		return nil, errors.Wrap(err, "Can't create new POST http.Request")
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	if len(c.username) > 0 {
		req.SetBasicAuth(c.username, c.password)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "Can't Do HTTP request by DefaultClient")
	}

	b, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, errors.Wrap(err, "Can't read bytes from HTTP response body")
	}

	return b, err
}

// GetError returns Ignite specific error message
func (c *client) GetError(successStatus SuccessStatus, error string) string {
	if successStatus < successStatusSuccess || successStatusSecurityCheckFailed < successStatus {
		successStatus = successStatusUnknown
	}
	m := strings.Join([]string{"Ignite returns: ", successStatusMsg[successStatus]}, "")
	if len(error) > 0 {
		m = strings.Join([]string{m, error}, ": ")
	}
	return m
}

// IsFailed returns `true` if `successStatus` value means failed
func (c *client) IsFailed(successStatus SuccessStatus) bool {
	return successStatus != successStatusSuccess
}
