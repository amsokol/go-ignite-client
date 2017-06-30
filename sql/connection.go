package sql

import (
	"context"
	"database/sql/driver"
)

// Conn extends driver.Conn
type Conn interface {
	driver.Conn

	CloseQueryContext(ctx context.Context, queryID int64) error

	ExecContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error)

	FetchContext(ctx context.Context, queryID int64, columns []Column) (*ResultSet, error)

	QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error)
}
