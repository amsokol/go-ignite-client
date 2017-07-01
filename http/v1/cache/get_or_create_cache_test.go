package cache

import (
	"testing"

	"github.com/amsokol/go-ignite-client/http/v1/exec"
)

func TestCommands_GetOrCreateCache(t *testing.T) {
	type args struct {
		e     exec.Executer
		cache string
	}
	tests := []struct {
		name      string
		c         *Commands
		args      args
		wantToken string
		wantErr   bool
	}{
		{
			name: "Create new cache",
			c:    &Commands{},
			args: args{
				e:     &exec.ExecuterImpl{Servers: []string{"http://localhost:8080/ignite"}, Username: "", Password: ""},
				cache: "Cache4TestGetOrCreateCache",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.c.GetOrCreateCache(tt.args.e, tt.args.cache)
			if (err != nil) != tt.wantErr {
				t.Errorf("Commands.GetOrCreateCache() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
