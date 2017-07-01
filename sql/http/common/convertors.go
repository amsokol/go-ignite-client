package common

import (
	"database/sql/driver"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/amsokol/go-ignite-client/sql"
)

// ItemsToValues converts [][]interface{} to [][]driver.Value
func ItemsToValues(columns []sql.Column, items [][]interface{}) ([][]driver.Value, error) {
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
			var i64 int64
			var f64 float64
			sv := fmt.Sprint(v)
			t := columns[j].ServerType
			switch t {
			case "java.lang.Byte":
				i64, err = strconv.ParseInt(sv, 10, 8)
				if err == nil {
					row[j] = int8(i64)
				}
			case "java.lang.Short":
				i64, err = strconv.ParseInt(sv, 10, 16)
				if err == nil {
					row[j] = int16(i64)
				}
			case "java.lang.Integer":
				i64, err = strconv.ParseInt(sv, 10, 32)
				if err == nil {
					row[j] = int32(i64)
				}
			case "java.lang.Long":
				row[j], err = strconv.ParseInt(sv, 10, 64)
			case "java.lang.Double":
				row[j], err = strconv.ParseFloat(sv, 64)
			case "java.lang.Float":
				f64, err = strconv.ParseFloat(sv, 64)
				if err == nil {
					row[j] = float32(f64)
				}
			case "java.lang.Boolean":
				row[j], err = strconv.ParseBool(sv)
			case "java.lang.Character":
				row[j] = []rune(sv)
			case "java.lang.String":
				row[j] = sv
			case "java.sql.Timestamp":
				// RFC1123 = "Mon, 02 Jan 2006 15:04:05 MST"
				// Custom:
				row[j], err = time.Parse("Jan 2, 2006 3:04:05 PM", sv)
			// TODO: add binary support
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

// NamedValuesToURLValues converts SQL parameters to HTTP request parameters
func NamedValuesToURLValues(nvs []driver.NamedValue) (url.Values, error) {
	vs := url.Values{}

	l := len(nvs)
	for i := 1; i <= l; i++ {
		for _, nv := range nvs {
			if nv.Ordinal == i {
				if nv.Value == nil {
					return nil, errors.New("Ignite HTTP REST API v1.x.x does not support NULL as parameter")
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
				case []rune:
					av = string(v)
				case string:
					av = v
				case time.Time:
					// RFC3339 = "2006-01-02T15:04:05Z07:00"
					av = v.Format(time.RFC3339)
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
