package v1

import (
	"github.com/pkg/errors"
)

// SQL result struct
type result struct {
}

// See https://golang.org/pkg/database/sql/driver/#Result for more details
func (r *result) LastInsertId() (int64, error) {
	return -1, errors.New("Ignite HTTP REST API v1.x.x does not support LastInsertId")
}

// See https://golang.org/pkg/database/sql/driver/#Result for more details
func (r *result) RowsAffected() (int64, error) {
	return -1, errors.New("Ignite HTTP REST API v1.x.x does not support RowsAffected")
}
