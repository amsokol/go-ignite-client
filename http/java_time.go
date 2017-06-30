package http

import (
	"strconv"
	"time"
	"unsafe"
)

// JavaTime is time in Java format
type JavaTime struct {
	time.Time
}

// UnmarshalJSON is unmarshaler for JavaTime
func (t *JavaTime) UnmarshalJSON(data []byte) error {
	m, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	gt := time.Unix(0, m*int64(time.Millisecond))
	*t = *(*JavaTime)(unsafe.Pointer(&gt))
	return nil
}
