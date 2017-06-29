package v10

import (
	"log"
	"testing"

	"github.com/amsokol/go-ignite-client/http/types"
	"github.com/amsokol/go-ignite-client/http/v1/client"
)

func TestDecrement(t *testing.T) {
	type args struct {
		c         client.Client
		cacheName string
		key       string
		init      *int64
		delta     int64
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		want1   string
		want2   types.SessionToken
		wantErr bool
	}{
		{
			name: "Decrement atomic long",
			args: args{
				c:         client.Open([]string{"http://localhost:8080/ignite"}, "", ""),
				cacheName: "Person",
				key:       "sequence",
				init:      nil,
				delta:     1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, err := Decrement(tt.args.c, tt.args.cacheName, tt.args.key, tt.args.init, tt.args.delta)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrement() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			log.Println("")
			log.Println("http.v1.v10.Decrement returned:")
			log.Println("value=", got)
			log.Println("affinityNodeId=", got1)
			log.Println("sessionToken=", got2)
		})
	}
}
