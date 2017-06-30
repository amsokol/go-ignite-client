package http

import (
	"encoding/json"
	"testing"
	"time"
)

func TestJavaTime_UnmarshalJSON(t *testing.T) {
	type JavaTimeTest struct {
		Data *JavaTime `json:"data"`
	}
	jtt := JavaTimeTest{}
	data := `{"data" : 1415179251551}`

	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		t       *JavaTimeTest
		args    args
		wantErr bool
	}{
		{
			name: "Parse date 1415179251551",
			t:    &jtt,
			args: args{
				data: []byte(data),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := json.Unmarshal(tt.args.data, tt.t); (err != nil) != tt.wantErr {
				t.Errorf("json.Unmarshal([]byte, JavaTimeTest) error = %v, wantErr %v", err, tt.wantErr)
			}
			if jtt.Data.Time.Year() != 2014 || jtt.Data.Time.Month() != time.November || jtt.Data.Time.Day() != 5 {
				t.Errorf("json.Unmarshal([]byte, JavaTimeTest) returns invalid result = %v", jtt.Data.Time)
			}
			t.Log("")
			t.Logf("http.v1.JavaTime.UnmarshalJSON returned for '%s':", tt.name)
			t.Log("value =", jtt.Data.Time)
		})
	}
}
