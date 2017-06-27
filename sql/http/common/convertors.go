package common

import (
	"database/sql/driver"
	"fmt"
	"net/url"
	"strconv"
	"strings"

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
