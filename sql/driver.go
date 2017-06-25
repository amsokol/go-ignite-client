package sql

import (
	"database/sql"

	"github.com/amsokol/go-ignite-client/sql/provider/http"
)

func init() {
	sql.Register("ignite-sql-http", &http.Driver{})
}
