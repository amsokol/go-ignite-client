package http

import (
	"context"
	"database/sql/driver"
)

// SQL statement struct
type stmt struct {
	connection *conn
	query      string
}

// See https://golang.org/pkg/database/sql/driver/#Stmt for more details
func (s *stmt) Close() error {
	s.connection = nil
	s.query = ""
	return nil
}

// See https://golang.org/pkg/database/sql/driver/#Stmt for more details
func (s *stmt) NumInput() int {
	return -1
}

// See https://golang.org/pkg/database/sql/driver/#Stmt for more details
func (s *stmt) Exec(args []driver.Value) (driver.Result, error) {
	return s.ExecContext(context.Background(), s.values2NamedValues(args))
}

// See https://golang.org/pkg/database/sql/driver/#StmtExecContext for more details
func (s *stmt) ExecContext(ctx context.Context, args []driver.NamedValue) (driver.Result, error) {
	return s.connection.ExecContext(ctx, s.query, args)
}

// See https://golang.org/pkg/database/sql/driver/#Stmt for more details
func (s *stmt) Query(args []driver.Value) (driver.Rows, error) {
	return s.QueryContext(context.Background(), s.values2NamedValues(args))
}

// See https://golang.org/pkg/database/sql/driver/#StmtQueryContext for more details
func (s *stmt) QueryContext(ctx context.Context, args []driver.NamedValue) (driver.Rows, error) {
	return s.connection.QueryContext(ctx, s.query, args)
}

// values2NamedValues converts []driver.Value slice to []driver.NamedValue slice
func (s *stmt) values2NamedValues(v []driver.Value) []driver.NamedValue {
	c := len(v)
	nv := make([]driver.NamedValue, c, c)
	for i, val := range v {
		nv[i] = driver.NamedValue{Value: val}
	}
	return nv
}
