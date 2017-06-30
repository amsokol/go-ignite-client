package sql

import (
	"context"
	"database/sql/driver"
	"io"

	"github.com/pkg/errors"
)

// ResultSet is SQL resultSet struct
type ResultSet struct {
	Last  bool
	Data  [][]driver.Value
	Index int
}

// Column is SQL column struct
type Column struct {
	Name       string
	ServerType string
}

// Rows is SQL rows struct
type Rows struct {
	Connection Conn
	QueryID    int64
	ColumnsRaw []Column
	ResultSet  *ResultSet
}

// GetResultAndMoveNext returns result with index current value and increase index by 1.
// Note: GetResultAndMoveNext does not check index out of range
func (rs *ResultSet) GetResultAndMoveNext() []driver.Value {
	r := rs.Data[rs.Index]
	rs.Index++
	return r
}

// Columns - see https://golang.org/pkg/database/sql/driver/#Rows for more details
func (r *Rows) Columns() []string {
	cl := len(r.ColumnsRaw)
	cs := make([]string, cl, cl)
	for i, v := range r.ColumnsRaw {
		cs[i] = v.Name
	}
	return cs
}

// ColumnTypeDatabaseTypeName - see https://golang.org/pkg/database/sql/driver/#RowsColumnTypeDatabaseTypeName for more details
func (r *Rows) ColumnTypeDatabaseTypeName(index int) string {
	if 0 > index || index >= len(r.ColumnsRaw) {
		return ""
	}
	return r.ColumnsRaw[index].ServerType
}

// Close - see https://golang.org/pkg/database/sql/driver/#Rows for more details
func (r *Rows) Close() error {
	defer func() {
		r.Connection = nil
		r.QueryID = 0
	}()

	if r.QueryID > 0 && r.Connection != nil {
		return r.Connection.CloseQueryContext(context.Background(), r.QueryID)
	}

	return nil
}

// Next - see https://golang.org/pkg/database/sql/driver/#Rows for more details
func (r *Rows) Next(dest []driver.Value) error {
	if r.Connection == nil {
		return errors.New("Rows are closed")
	}

	if r.ResultSet == nil {
		return io.EOF
	}

	size := len(r.ResultSet.Data)
	if r.ResultSet.Index >= size {
		if r.ResultSet.Last {
			return io.EOF
		}

		var err error
		r.ResultSet, err = r.Connection.FetchContext(context.Background(), r.QueryID, r.ColumnsRaw)
		if err != nil {
			return errors.Wrap(err, "Failed to get next page for the query")
		}

		if len(r.ResultSet.Data) == 0 {
			return io.EOF
		}
	}

	row := r.ResultSet.GetResultAndMoveNext()
	for i, v := range row {
		dest[i] = v
	}

	return nil
}
