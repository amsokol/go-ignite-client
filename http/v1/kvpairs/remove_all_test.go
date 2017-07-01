package kvpairs

import (
	"net/url"
	"testing"

	"github.com/amsokol/go-ignite-client/http/v1/exec"
)

func TestCommands_RemoveAll(t *testing.T) {
	v := url.Values{}
	v.Add("k1", "1")
	v.Add("k2", "2")

	type args struct {
		e      exec.Executer
		cache  string
		args   url.Values
		destID string
	}
	tests := []struct {
		name       string
		c          *Commands
		args       args
		wantOk     bool
		wantNodeID string
		wantToken  string
		wantErr    bool
	}{
		{
			name: "Remove All",
			c:    &Commands{},
			args: args{
				e:     &exec.ExecuterImpl{Servers: []string{"http://localhost:8080/ignite"}, Username: "", Password: ""},
				cache: "Organization",
				args:  v,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOk, _, _, err := tt.c.RemoveAll(tt.args.e, tt.args.cache, tt.args.args, tt.args.destID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Commands.RemoveAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOk != tt.wantOk {
				t.Errorf("Commands.RemoveAll() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
