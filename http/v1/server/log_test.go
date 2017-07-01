package server

import (
	"testing"

	"github.com/amsokol/go-ignite-client/http/v1/exec"
)

func TestCommands_GetLog(t *testing.T) {
	from := 10
	to := 15

	type args struct {
		e    exec.Executer
		path string
		from *int
		to   *int
	}
	tests := []struct {
		name      string
		c         *Commands
		args      args
		wantLog   string
		wantToken string
		wantErr   bool
	}{
		{
			name: "Get log from 10 to 15 line",
			c:    &Commands{},
			args: args{
				e:    &exec.ExecuterImpl{Servers: []string{"http://localhost:8080/ignite"}, Username: "", Password: ""},
				from: &from,
				to:   &to,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLog, gotToken, err := tt.c.GetLog(tt.args.e, tt.args.path, tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("Commands.GetLog() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log("")
			t.Logf("http.v1.server.GetLog returned for '%s':", tt.name)
			t.Log("log =", gotLog)
			t.Log("sessionToken =", gotToken)
		})
	}
}
