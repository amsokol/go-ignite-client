package server

import (
	"testing"

	"github.com/amsokol/go-ignite-client/http/v1/exec"
)

func TestCommands_GetVersion(t *testing.T) {
	type args struct {
		e exec.Executer
	}
	tests := []struct {
		name        string
		c           *Commands
		args        args
		wantVersion string
		wantToken   string
		wantErr     bool
	}{
		{
			name: "Get version of server",
			c:    &Commands{},
			args: args{
				e: &exec.ExecuterImpl{Servers: []string{"http://localhost:8080/ignite"}, Username: "", Password: ""},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVersion, gotToken, err := tt.c.GetVersion(tt.args.e)
			if (err != nil) != tt.wantErr {
				t.Errorf("Commands.GetVersion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log("")
			t.Logf("http.v1.server.GetVersion returned for '%s':", tt.name)
			t.Log("version =", gotVersion)
			t.Log("sessionToken =", gotToken)
		})
	}
}
