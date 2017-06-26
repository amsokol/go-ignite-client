package http

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/blang/semver"
	"github.com/pkg/errors"

	"github.com/amsokol/go-ignite-client/http/internal"
	"io"
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
func Version(servers []string, quarantine float64, username string, password string) (*semver.Version, error) {
	v := url.Values{}
	v.Add("cmd", "version")

	var server string
	server = ""
	for {
		server, err := internal.GlobalPool.GetNextServer(server, servers, quarantine)
		if err != nil {
			if err == io.EOF {
				return nil, errors.Wrap(err, "All servers are down or not available for you")
			}
			return nil, errors.Wrap(err, "Can't get server from pool")
		}

		req, err := http.NewRequest("POST", server, strings.NewReader(v.Encode()))
		if err != nil {
			return nil, errors.Wrap(err, "Can't create new POST http.Request")
		}

		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		if len(username) > 0 {
			req.SetBasicAuth(username, password)
		}

		// log.Println("VERSION POSTing request to server", server)

		r, err := http.DefaultClient.Do(req)
		if err == nil {
			if !isServerDown(r.StatusCode) {
				internal.GlobalPool.UpdateStatus(server, true)

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
				return &ver, nil
			}
		}
		internal.GlobalPool.UpdateStatus(server, false)
		log.Println("Server", server, "is down or not available for you:", err)
	}
}

func isServerDown(code int) bool {
	return code == http.StatusBadGateway || code == http.StatusInternalServerError
}
