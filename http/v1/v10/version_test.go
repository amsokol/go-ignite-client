package v10

import (
	"log"
	"testing"

	"github.com/amsokol/go-ignite-client/http/types"
	"github.com/amsokol/go-ignite-client/http/v1/client"
)

func TestVersion(t *testing.T) {
	type args struct {
		c client.Client
	}
	tests := []struct {
		name    string
		args    args
		want    types.Version
		want1   types.SessionToken
		wantErr bool
	}{
		{
			name: "Get log from 10 to 15 line",
			args: args{
				c: client.Open([]string{"http://localhost:8080/ignite"}, "", ""),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := Version(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("Version() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			log.Println("")
			log.Println("http.v1.v10.Version returned:")
			log.Println("version=", got)
			log.Println("sessionToken=", got1)
		})
	}
}
