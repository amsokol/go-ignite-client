package internal

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"

	"github.com/amsokol/go-ignite-client/http/internal"
	"github.com/amsokol/go-ignite-client/http/types"
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

// Client is the object providing the methods to execute REST API commands
type Client struct {
	Servers  []string
	Username string
	Password string
}

// Execute implements http.CommandExecutor
func (c *Client) Execute(v url.Values) ([]byte, error) {
	var server string
	server = ""
	for i := 0; i < len(c.Servers); i++ {
		server, err := internal.GlobalPool.GetNextServer(server, c.Servers)
		if err != nil {
			return nil, errors.Wrap(err, "Can't get server from pool")
		}

		req, err := http.NewRequest("POST", server, strings.NewReader(v.Encode()))
		if err != nil {
			return nil, errors.Wrap(err, "Can't create new POST http.Request")
		}

		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		if len(c.Username) > 0 {
			req.SetBasicAuth(c.Username, c.Password)
		}

		//		log.Println("CLIENT POSTing request to server", server)

		res, err := http.DefaultClient.Do(req)
		if err == nil {
			if !c.isServerDown(res.StatusCode) {
				internal.GlobalPool.UpdateStatus(server, true)

				b, err := ioutil.ReadAll(res.Body)
				res.Body.Close()
				if err != nil {
					return nil, errors.Wrap(err, "Can't read bytes from HTTP response body")
				}

				return b, nil
			}
		}
		internal.GlobalPool.UpdateStatus(server, false)
		log.Println("Server", server, "is down or not available for you:", err)
	}
	return nil, errors.New("All servers are down or not available for you")
}

// GetError returns Ignite specific error message
func (c *Client) GetError(successStatus types.SuccessStatus, error string) string {
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
func (c *Client) IsFailed(successStatus types.SuccessStatus) bool {
	return successStatus != successStatusSuccess
}

func (c *Client) isServerDown(code int) bool {
	return code == http.StatusBadGateway || code == http.StatusInternalServerError
}
