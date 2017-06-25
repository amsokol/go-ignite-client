package http

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/pkg/errors"

	"github.com/amsokol/go-ignite-client/http"
)

// Driver is exported to allow it to be used directly.
type Driver struct{}

// Open a Connection to the server.
func (a *Driver) Open(name string) (driver.Conn, error) {
	ci := http.ConnectionInfo{}

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

	c := &http.Client{ConnectionInfo: &ci}

	return &conn{client: c}, nil
}
