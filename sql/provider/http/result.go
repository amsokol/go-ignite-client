package http

import (
	"github.com/pkg/errors"
)

type result struct {
	lastInsertId int64
	rowsAffected int64
}

func (r *result) LastInsertId() (int64, error) {
	return -1, errors.WithStack(errors.New("Ignite HTTP REST API does not support LastInsertId"))
}

func (r *result) RowsAffected() (int64, error) {
	return -1, errors.WithStack(errors.New("Ignite HTTP REST API does not support RowsAffected"))
}
