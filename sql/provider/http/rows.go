package http

import (
	"context"
	"database/sql/driver"
	"io"

	"github.com/pkg/errors"
)

// SQL column struct
type column struct {
	name       string
	serverType string
}

// SQL rows struct
type rows struct {
	connection *conn
	queryID    string
	columns    []column
	resultSet  *resultSet
}

// See https://golang.org/pkg/database/sql/driver/#Rows for more details
func (r *rows) Columns() []string {
	cl := len(r.columns)
	cs := make([]string, cl, cl)
	for i, v := range r.columns {
		cs[i] = v.name
	}
	return cs
}

// See https://golang.org/pkg/database/sql/driver/#RowsColumnTypeDatabaseTypeName for more details
func (r *rows) ColumnTypeDatabaseTypeName(index int) string {
	if 0 > index || index >= len(r.columns) {
		return ""
	}
	return r.columns[index].serverType
}

// See https://golang.org/pkg/database/sql/driver/#Rows for more details
func (r *rows) Close() error {
	defer func() {
		r.connection = nil
		r.queryID = ""
	}()

	if len(r.queryID) > 0 && r.connection != nil {
		return r.connection.closeQueryContext(context.Background(), r.queryID)
	}

	return nil
}

// See https://golang.org/pkg/database/sql/driver/#Rows for more details
func (r *rows) Next(dest []driver.Value) error {
	if r.connection == nil {
		return errors.New("Rows are closed")
	}

	if r.resultSet == nil {
		return io.EOF
	}

	size := len(r.resultSet.data)
	if r.resultSet.index >= size {
		if r.resultSet.last {
			return io.EOF
		}

		var err error
		r.resultSet, err = r.connection.fetchContext(context.Background(), r.queryID, r.columns)
		if err != nil {
			return errors.Wrap(err, "Failed to get next page for the query")
		}

		if len(r.resultSet.data) == 0 {
			return io.EOF
		}
	}

	row := r.resultSet.getResultAndMoveNext()
	for i, v := range row {
		dest[i] = v
	}

	return nil
}
