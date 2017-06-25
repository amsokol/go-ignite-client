package http

import (
	"context"
	"database/sql/driver"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// SQL column struct
type column struct {
	name       string
	igniteType string
}

// SQL rows struct
type rows struct {
	connection *conn
	queryID    string
	last       bool
	columns    []column
	resultSet  [][]driver.Value
	index      int
	size       int
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
	return r.columns[index].igniteType
}

// See https://golang.org/pkg/database/sql/driver/#Rows for more details
func (r *rows) Close() (err error) {
	if r.connection != nil && len(r.queryID) > 0 {
		err = r.connection.closeQueryContext(context.Background(), r.queryID)
	}

	r.connection = nil
	r.queryID = ""
	return err
}

// See https://golang.org/pkg/database/sql/driver/#Rows for more details
func (r *rows) Next(dest []driver.Value) error {
	if r.connection == nil {
		return errors.New("Rows are closed")
	}

	if r.index >= r.size {
		if r.last {
			return io.EOF
		}
		items, last, err := r.connection.fetchContext(context.Background(), r.queryID)
		if err != nil {
			return errors.Wrap(err, "Failed to get next page for the query")
		}

		r.last = last
		r.setResultSet(items)
		if err != nil {
			return errors.Wrap(err, "Failed extract ResultSet from fetch response")
		}

		if r.size == 0 {
			r.last = true
			return io.EOF
		}
	}

	row := r.resultSet[r.index]
	for i, v := range row {
		dest[i] = v
	}
	r.index++

	return nil
}

// setResultSet sets rows result set
func (r *rows) setResultSet(items [][]interface{}) error {
	il := len(items)
	rs := make([][]driver.Value, il, il)

	cl := len(r.columns)
	for i, item := range items {
		if cl != len(item) {
			return errors.New("Very strange situation - column count and count of values in row are different")
		}
		row := make([]driver.Value, cl, cl)
		for j, v := range item {
			var err error
			sv := fmt.Sprint(v)
			t := r.columns[j].igniteType
			switch t {
			case "java.lang.Byte":
				row[j], err = strconv.ParseInt(sv, 10, 8)
			case "java.lang.Short":
				row[j], err = strconv.ParseInt(sv, 10, 16)
			case "java.lang.Integer":
				row[j], err = strconv.ParseInt(sv, 10, 32)
			case "java.lang.Long":
				row[j], err = strconv.ParseInt(sv, 10, 64)
			case "java.lang.Double":
				row[j], err = strconv.ParseFloat(sv, 64)
			case "java.lang.Boolean":
				row[j], err = strconv.ParseBool(sv)
			case "java.lang.Character":
				row[j] = sv
			case "java.lang.String":
				row[j] = sv
			// TODO: add binary support
			// TODO: add time.Time support
			default:
				return errors.New(strings.Join([]string{"Unsupported parameter type", t}, ": "))
			}
			if err != nil {
				return errors.Wrap(err, strings.Join([]string{"Failed to convert Ignite type to golang type", t}, ": "))
			}
		}
		rs[i] = row
	}

	r.resultSet = rs
	r.size = il
	r.index = 0
	return nil
}
