package v2

import (
	"context"
	"database/sql/driver"
	"strconv"

	"github.com/pkg/errors"

	"github.com/amsokol/go-ignite-client/http/v2"
	"github.com/amsokol/go-ignite-client/sql"
	"github.com/amsokol/go-ignite-client/sql/http/common"
)

// SQL connection struct
type conn struct {
	client     v2.Client
	cache      string
	pageSize   int64
	quarantine float64
}

// Open returns connection
func Open(servers []string, username string, password string, cache string, pageSize int64) sql.Conn {
	return &conn{cache: cache, pageSize: pageSize, client: v2.Open(servers, username, password)}
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

	v, err := common.NamedValuesToURLValues(args)
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

	v, err := common.NamedValuesToURLValues(args)
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
	data, err := common.ItemsToValues(columns, res.GetItems())
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
	data, err := common.ItemsToValues(columns, res.GetItems())
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

func (c *conn) getResultSet(data [][]driver.Value, last bool) *sql.ResultSet {
	if len(data) == 0 {
		last = true
	}
	return &sql.ResultSet{Data: data, Index: 0, Last: last}
}
