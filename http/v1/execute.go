package v1

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"

	"github.com/amsokol/go-ignite-client/http/internal"
	"strconv"
)

// See https://apacheignite.readme.io/docs/rest-api#section-returned-value for more details
const (
	successStatusSuccess             = 0
	successStatusFailed              = 1
	successStatusAuthorizationFailed = 2
	successStatusSecurityCheckFailed = 3
)

type responseWrapper struct {
	SuccessStatus  int              `json:"successStatus"`
	Error          string           `json:"error"`
	SessionToken   string           `json:"sessionToken"`
	AffinityNodeID string           `json:"affinityNodeId"`
	Response       abstractResponse `json:"response"`
}

type abstractResponse struct {
	Data []byte
}

// UnmarshalJSON is unmarshaler for abstractResponse
func (r *abstractResponse) UnmarshalJSON(data []byte) error {
	r.Data = data
	return nil
}

// Execute implements http.CommandExecutor
func (c *client) execute(v url.Values) (response []byte, nodeID string, token string, err error) {
	resWrapper := responseWrapper{}

	var server string
	server = ""
	for i := 0; i < len(c.servers); i++ {
		server, err := internal.GlobalPool.GetNextServer(server, c.servers)
		if err != nil {
			return resWrapper.Response.Data, resWrapper.AffinityNodeID, resWrapper.SessionToken, errors.Wrap(err, "Can't get server from pool")
		}

		req, err := http.NewRequest("POST", server, strings.NewReader(v.Encode()))
		if err != nil {
			return resWrapper.Response.Data, resWrapper.AffinityNodeID, resWrapper.SessionToken, errors.Wrap(err, "Can't create new POST http.Request")
		}

		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		if len(c.username) > 0 {
			req.SetBasicAuth(c.username, c.password)
		}

		//		log.Println("CLIENT POSTing request to server", server)

		res, err := http.DefaultClient.Do(req)
		if err == nil {
			if res.StatusCode != http.StatusBadGateway && res.StatusCode != http.StatusInternalServerError {
				internal.GlobalPool.UpdateStatus(server, true)

				b, err := ioutil.ReadAll(res.Body)
				res.Body.Close()
				if err != nil {
					return resWrapper.Response.Data, resWrapper.AffinityNodeID, resWrapper.SessionToken, errors.Wrap(err, "Can't read bytes from HTTP response body")
				}

				// log.Println("Raw response:", string(b))

				err = json.Unmarshal(b, &resWrapper)
				if err != nil {
					return resWrapper.Response.Data, resWrapper.AffinityNodeID, resWrapper.SessionToken, errors.Wrap(err, "Can't unmarshal respone to response struct")
				}

				if resWrapper.SuccessStatus != successStatusSuccess {
					return resWrapper.Response.Data, resWrapper.AffinityNodeID, resWrapper.SessionToken,
						errors.New(strings.Join([]string{"successStatus=",
							strconv.FormatInt(int64(resWrapper.SuccessStatus), 10), ", error=", resWrapper.Error}, ""))
				}

				// log.Println("Response:", res)

				return resWrapper.Response.Data, resWrapper.AffinityNodeID, resWrapper.SessionToken, nil
			}
		}
		internal.GlobalPool.UpdateStatus(server, false)
		log.Println("Server", server, "is down or not available for you:", err)
	}
	return resWrapper.Response.Data, resWrapper.AffinityNodeID, resWrapper.SessionToken, errors.New("All servers are down or not available for you")
}
