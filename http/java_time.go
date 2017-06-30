package http

import (
	"strconv"
	"time"
)

// JavaTime is time in Java format
type JavaTime struct {
	Time time.Time
}

// UnmarshalJSON is unmarshaler for JavaTime
func (t *JavaTime) UnmarshalJSON(data []byte) error {
	m, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	*t = JavaTime{Time: time.Unix(0, m*int64(time.Millisecond))}
	return nil
}
