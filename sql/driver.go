package sql

import (
	"database/sql"

	"github.com/amsokol/go-ignite-client/sql/provider/http"
)

// Initialize driver
func init() {
	sql.Register("ignite-sql-http", &http.Driver{})
}
