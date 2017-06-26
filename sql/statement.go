package sql

import (
	"context"
	"database/sql/driver"

	"github.com/pkg/errors"
)

// Stmt is SQL statement struct
type Stmt struct {
	Connection Conn
	QueryExp   string
}

// Close - see https://golang.org/pkg/database/sql/driver/#Stmt for more details
func (s *Stmt) Close() error {
	s.Connection = nil
	return nil
}

// NumInput - see https://golang.org/pkg/database/sql/driver/#Stmt for more details
func (s *Stmt) NumInput() int {
	return -1
}

// Exec - see https://golang.org/pkg/database/sql/driver/#Stmt for more details
func (s *Stmt) Exec(args []driver.Value) (driver.Result, error) {
	return s.ExecContext(context.Background(), s.values2NamedValues(args))
}

// ExecContext - see https://golang.org/pkg/database/sql/driver/#StmtExecContext for more details
func (s *Stmt) ExecContext(ctx context.Context, args []driver.NamedValue) (driver.Result, error) {
	if s.Connection == nil {
		return nil, errors.New("Statement is closed")
	}
	return s.Connection.ExecContext(ctx, s.QueryExp, args)
}

// Query - see https://golang.org/pkg/database/sql/driver/#Stmt for more details
func (s *Stmt) Query(args []driver.Value) (driver.Rows, error) {
	return s.QueryContext(context.Background(), s.values2NamedValues(args))
}

// QueryContext - see https://golang.org/pkg/database/sql/driver/#StmtQueryContext for more details
func (s *Stmt) QueryContext(ctx context.Context, args []driver.NamedValue) (driver.Rows, error) {
	if s.Connection == nil {
		return nil, errors.New("Statement is closed")
	}
	return s.Connection.QueryContext(ctx, s.QueryExp, args)
}

// values2NamedValues converts []driver.Value slice to []driver.NamedValue slice
func (s *Stmt) values2NamedValues(v []driver.Value) []driver.NamedValue {
	c := len(v)
	nv := make([]driver.NamedValue, c, c)
	for i, val := range v {
		nv[i] = driver.NamedValue{Value: val}
	}
	return nv
}
