package v10

import (
	"log"
	"testing"

	"github.com/amsokol/go-ignite-client/http/types"
	"github.com/amsokol/go-ignite-client/http/v1/client"
)

func TestLog(t *testing.T) {
	type args struct {
		c    client.Client
		path string
		from int
		to   int
	}

	tests := []struct {
		name    string
		args    args
		want    string
		want1   types.SessionToken
		wantErr bool
	}{
		{
			name: "Get log from 10 to 15 line",
			args: args{
				c:    client.Open([]string{"http://localhost:8080/ignite"}, "", ""),
				from: 10,
				to:   15,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2, err := Log(tt.args.c, tt.args.path, tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("Log() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			log.Println("")
			log.Println("http.v1.v10.Log returned:")
			log.Println("log=", got1)
			log.Println("sessionToken=", got2)
		})
	}
}
