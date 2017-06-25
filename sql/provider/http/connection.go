package http

import (
	"context"
	"database/sql/driver"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/amsokol/go-ignite-client/http"
)

// SQL connection struct
type conn struct {
	client *http.Client
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
	return &stmt{connection: c, query: query}, nil
}

// See https://golang.org/pkg/database/sql/driver/#Conn for more details
func (c *conn) Close() error {
	c.client = nil

	return nil
}

// See https://golang.org/pkg/database/sql/driver/#ConnBeginTx for more details
func (c *conn) BeginTx(ctx context.Context, opts driver.TxOptions) (driver.Tx, error) {
	return nil, errors.New("Ignite REST API does not support transactions")
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

	_, _, err = c.client.QryFldExe(query, v)
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

	_, _, err := c.client.Version()
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

	r, _, err := c.client.QryFldExe(query, v)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to invoke 'qryfldexe' command")
	}

	cl := len(r.FieldsMetadata)
	rows := rows{connection: c, queryID: fmt.Sprintf("%d", r.QueryID), last: r.Last, columns: make([]column, cl, cl)}

	// columns
	for i, c := range r.FieldsMetadata {
		rows.columns[i] = column{name: c.FieldName, igniteType: c.FieldTypeName}
	}

	// data
	err = rows.setResultSet(r.Items)
	if err != nil {
		return nil, errors.Wrap(err, "Failed extract ResultSet from 'qryfldexe' response")
	}

	return &rows, nil
}

// fetchContext gets next page for the query
func (c *conn) fetchContext(ctx context.Context, queryID string) ([][]interface{}, bool, error) {
	if c.client == nil {
		return nil, false, driver.ErrBadConn
	}

	r, _, err := c.client.QryFetch(queryID)
	if err != nil {
		return nil, false, errors.Wrap(err, "Failed to invoke 'qryfldexe' command")
	}

	return r.Items, r.Last, nil
}

// closeQueryContext closes query resources
func (c *conn) closeQueryContext(ctx context.Context, queryID string) error {
	if c.client == nil {
		return driver.ErrBadConn
	}

	_, _, err := c.client.QryCls(queryID)

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
					return nil, errors.WithStack(errors.New("Ignite HTTP REST API does not support NULL as parameter"))
				}
				var av string
				switch v := nv.Value.(type) {
				case int8:
					av = fmt.Sprintf("%d", v)
				case int16:
					av = fmt.Sprintf("%d", v)
				case int32:
					av = fmt.Sprintf("%d", v)
				case int64:
					av = fmt.Sprintf("%d", v)
				case float64:
					av = fmt.Sprintf("%f", v)
				case float32:
					av = fmt.Sprintf("%f", v)
				case bool:
					av = fmt.Sprintf("%t", v)
				case string:
					av = v
				// TODO: add binary support
				// TODO: add time.Time support
				default:
					return nil, errors.WithStack(errors.New(strings.Join([]string{"Unsupported parameter type with index", strconv.Itoa(i)}, " ")))
				}
				vs.Add(strings.Join([]string{"arg", strconv.Itoa(i)}, ""), av)
				break
			}
		}
	}
	return vs, nil
}
