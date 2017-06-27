package http

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/amsokol/go-ignite-client/sql/http/v1"
	"github.com/amsokol/go-ignite-client/sql/http/v2"
)

type connectionInfo struct {
	Version    float64  `json:"version"`
	Servers    []string `json:"servers"`
	Username   string   `json:"username"`
	Password   string   `json:"password"`
	Cache      string   `json:"cache"`
	PageSize   int64    `json:"pageSize"`
	Quarantine float64  `json:"quarantine"`
}

// Driver is exported to allow it to be used directly.
type Driver struct{}

// Open a Connection to the server.
func (a *Driver) Open(name string) (driver.Conn, error) {
	ci := connectionInfo{}

	err := json.Unmarshal([]byte(name), &ci)
	if err != nil {
		return nil, errors.Wrap(err, "Invalid format of connection information")
	}

	if len(ci.Servers) == 0 {
		return nil, errors.New("You need to provide at least one server")
	}

	if len(ci.Cache) == 0 {
		return nil, errors.New("You need to provide cache name")
	}

	if ci.PageSize == 0 {
		// set default value 1000
		ci.PageSize = 1000
	}

	if ci.Quarantine == 0 {
		// set default value 30 min
		ci.Quarantine = 30
	}

	if ci.Version == 0 {
		// set default version is 1.0
		ci.Version = 1.0
	}

	var conn driver.Conn
	switch ci.Version {
	case 1:
		conn = v1.Open(ci.Servers, ci.Quarantine, ci.Username, ci.Password, ci.Cache, ci.PageSize)
	case 2:
		conn = v2.Open(ci.Servers, ci.Quarantine, ci.Username, ci.Password, ci.Cache, ci.PageSize)
	default:
		return nil, errors.New(strings.Join([]string{"Unsupported HTTP REST API version v",
			strconv.FormatFloat(ci.Version, 'f', -1, 64), ". Supported versions are \"1\" and \"2\""}, ""))
	}

	return conn, nil
}

// Init Initializes driver
func init() {
	sql.Register("ignite-sql-http", &Driver{})
}
