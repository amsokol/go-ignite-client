package v2

import (
	"context"
	"database/sql/driver"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/amsokol/go-ignite-client/http/v2"
	"github.com/amsokol/go-ignite-client/sql"
)

// SQL connection struct
type conn struct {
	client     v2.Client
	cache      string
	pageSize   int64
	quarantine float64
}

// Open returns connection
func Open(servers []string, quarantine float64, username string, password string, cache string, pageSize int64) sql.Conn {
	return &conn{cache: cache, pageSize: pageSize, client: v2.Open(servers, quarantine, username, password)}
}

// See https://golang.org/pkg/database/sql/driver/#Conn for more details
func (c *conn) Prepare(query string) (driver.Stmt, error) {
	return c.PrepareContext(context.Background(), query)
}

// See https://golang.org/pkg/database/sql/driver/#ConnPrepareContext for more details
func (c *conn) PrepareContext(ctx context.Context, query string) (driver.Stmt, error) {
	if c.client == nil {
		return nil, driver.ErrBadConn
	}
	return &sql.Stmt{Connection: c, QueryExp: query}, nil
}

// See https://golang.org/pkg/database/sql/driver/#Conn for more details
func (c *conn) Close() error {
	c.client = nil

	return nil
}

// See https://golang.org/pkg/database/sql/driver/#ConnBeginTx for more details
func (c *conn) BeginTx(ctx context.Context, opts driver.TxOptions) (driver.Tx, error) {
	return nil, errors.New("Ignite REST API v2.x.x does not support transactions")
}

// See https://golang.org/pkg/database/sql/driver/#Conn
func (c *conn) Begin() (driver.Tx, error) {
	return c.BeginTx(nil, driver.TxOptions{})
}

// See https://golang.org/pkg/database/sql/driver/#StmtExecContext for more details
func (c *conn) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error) {
	if c.client == nil {
		return nil, driver.ErrBadConn
	}

	v, err := c.namedValues2UrlValues(args)
	if err != nil {
		return nil, errors.Wrap(err, "Failed convert parameters for REST API")
	}

	_, _, err = c.client.SQLFieldsQueryExecute(c.cache, c.pageSize, query, v)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to invoke 'qryfldexe' command")
	}

	return &result{}, nil
}

// See https://golang.org/pkg/database/sql/driver/#Pinger for more details
func (c *conn) Ping(ctx context.Context) error {
	if c.client == nil {
		return driver.ErrBadConn
	}

	_, _, err := c.client.GetVersion()
	return err
}

// See https://golang.org/pkg/database/sql/driver/#QueryerContext for more details
func (c *conn) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
	if c.client == nil {
		return nil, driver.ErrBadConn
	}

	v, err := c.namedValues2UrlValues(args)
	if err != nil {
		return nil, errors.Wrap(err, "Failed convert parameters for REST API 'qryfldexe' command")
	}

	res, _, err := c.client.SQLFieldsQueryExecute(c.cache, c.pageSize, query, v)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to invoke 'qryfldexe' command")
	}

	// columns
	colcount := len(res.GetFieldsMetadata())
	columns := make([]sql.Column, colcount, colcount)
	for i, c := range res.GetFieldsMetadata() {
		columns[i] = sql.Column{Name: c.GetFieldName(), ServerType: c.GetFieldTypeName()}
	}

	// data
	data, err := c.items2Values(columns, res.GetItems())
	if err != nil {
		return nil, errors.Wrap(err, "Failed extract values from 'qryfldexe' response")
	}

	return &sql.Rows{Connection: c,
		ColumnsRaw: columns,
		QueryID:    strconv.FormatInt(res.GetQueryID(), 10),
		ResultSet:  c.getResultSet(data, res.GetLast())}, nil
}

// fetchContext gets next page for the query
func (c *conn) FetchContext(ctx context.Context, queryID string, columns []sql.Column) (*sql.ResultSet, error) {
	if c.client == nil {
		return nil, driver.ErrBadConn
	}

	res, _, err := c.client.SQLQueryFetch(c.pageSize, queryID)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to invoke 'qryfetch' command")
	}

	// data
	data, err := c.items2Values(columns, res.GetItems())
	if err != nil {
		return nil, errors.Wrap(err, "Failed extract values from 'qryfetch' response")
	}

	return c.getResultSet(data, res.GetLast()), nil
}

// closeQueryContext closes query resources
func (c *conn) CloseQueryContext(ctx context.Context, queryID string) error {
	if c.client == nil {
		return driver.ErrBadConn
	}

	_, _, err := c.client.SQLQueryClose(queryID)

	return err
}

// namedValues2UrlValues converts SQL parameters to HTTP request parameters
func (c *conn) namedValues2UrlValues(nvs []driver.NamedValue) (url.Values, error) {
	vs := url.Values{}

	l := len(nvs)
	for i := 1; i <= l; i++ {
		for _, nv := range nvs {
			if nv.Ordinal == i {
				if nv.Value == nil {
					return nil, errors.New("Ignite HTTP REST API v2.x.x does not support NULL as parameter")
				}
				var av string
				switch v := nv.Value.(type) {
				case int8:
					av = strconv.FormatInt(int64(int8(v)), 10)
				case int16:
					av = strconv.FormatInt(int64(int16(v)), 10)
				case int32:
					av = strconv.FormatInt(int64(int32(v)), 10)
				case int64:
					av = strconv.FormatInt(int64(v), 10)
				case float64:
					av = strconv.FormatFloat(float64(v), 'f', -1, 64)
				case float32:
					av = strconv.FormatFloat(float64(float32(v)), 'f', -1, 32)
				case bool:
					av = strconv.FormatBool(bool(v))
				case string:
					av = v
				// TODO: add binary support
				// TODO: add time.Time support
				default:
					return nil, errors.New(strings.Join([]string{"Unsupported parameter type with index", strconv.Itoa(i)}, " "))
				}
				vs.Add(strings.Join([]string{"arg", strconv.Itoa(i)}, ""), av)
				break
			}
		}
	}
	return vs, nil
}

// setResultSet sets rows result set
func (c *conn) items2Values(columns []sql.Column, items [][]interface{}) ([][]driver.Value, error) {
	size := len(items)
	data := make([][]driver.Value, size, size)

	colcount := len(columns)
	for i, item := range items {
		if colcount != len(item) {
			return nil, errors.New("It's very strange situation - column count and count of values in row are different")
		}
		row := make([]driver.Value, colcount, colcount)
		for j, v := range item {
			var err error
			sv := fmt.Sprint(v)
			t := columns[j].ServerType
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
				return nil, errors.New(strings.Join([]string{"Unsupported parameter type", t}, ": "))
			}
			if err != nil {
				return nil, errors.Wrap(err, strings.Join([]string{"Failed to convert Ignite type to golang type", t}, ": "))
			}
		}
		data[i] = row
	}

	return data, nil
}

func (c *conn) getResultSet(data [][]driver.Value, last bool) *sql.ResultSet {
	if len(data) == 0 {
		last = true
	}
	return &sql.ResultSet{Data: data, Index: 0, Last: last}
}
