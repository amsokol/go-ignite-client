package http

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"strings"

	"github.com/pkg/errors"

	"github.com/amsokol/go-ignite-client/http"
	"github.com/amsokol/go-ignite-client/sql/http/v1"
)

type connectionInfo struct {
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

	ver, err := http.Version(ci.Servers, ci.Quarantine, ci.Username, ci.Password)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get HTTP REST API version")
	}

	var conn driver.Conn
	switch ver.Major {
	case 1:
		conn = v1.Open(ci.Servers, ci.Quarantine, ci.Username, ci.Password, ci.Cache, ci.PageSize)
	case 2:
		conn = v1.Open(ci.Servers, ci.Quarantine, ci.Username, ci.Password, ci.Cache, ci.PageSize)
	default:
		return nil, errors.Wrap(err, strings.Join([]string{"Unsupported HTTP REST API version v", ver.String()}, ""))
	}

	return conn, nil
}

// Init Initializes driver
func init() {
	sql.Register("ignite-sql-http", &Driver{})
}
