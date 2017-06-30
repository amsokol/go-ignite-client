package v1

import (
	"testing"
)

func Test_client_GetLog(t *testing.T) {
	from := 10
	to := 15

	type args struct {
		path string
		from *int
		to   *int
	}
	tests := []struct {
		name      string
		c         Client
		args      args
		wantLog   string
		wantToken string
		wantErr   bool
	}{
		{
			name: "Get log from 10 to 15 line",
			c:    NewClient([]string{"http://localhost:8080/ignite"}, "", ""),
			args: args{
				from: &from,
				to:   &to,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLog, gotToken, err := tt.c.GetLog(tt.args.path, tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.GetLog() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log("")
			t.Logf("http.v1.client.GetLog returned for '%s':", tt.name)
			t.Log("log =", gotLog)
			t.Log("sessionToken =", gotToken)
		})
	}
}
