package http

import (
	"database/sql/driver"
)

// SQL resultSet struct
type resultSet struct {
	last  bool
	data  [][]driver.Value
	index int
}

// getResult returns result with index current value and increase index by 1.
// Note: getResult does not check index out of range
func (rs *resultSet) getResultAndMoveNext() []driver.Value {
	r := rs.data[rs.index]
	rs.index++
	return r
}
