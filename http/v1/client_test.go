package v1

import (
	"reflect"
	"testing"
)

func Test_client_Close(t *testing.T) {
	tests := []struct {
		name    string
		c       Client
		wantErr bool
	}{
		{
			name: "Close client",
			c:    NewClient([]string{"http://localhost:8080/ignite"}, "", ""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.Close(); (err != nil) != tt.wantErr {
				t.Errorf("client.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewClient(t *testing.T) {
	type args struct {
		servers  []string
		username string
		password string
	}
	tests := []struct {
		name string
		args args
		want Client
	}{
		{
			name: "Create new Client",
			args: args{
				servers:  []string{"http://localhost:8080/ignite"},
				username: "",
				password: "",
			},
			want: &client{servers: []string{"http://localhost:8080/ignite"}, username: "", password: ""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(tt.args.servers, tt.args.username, tt.args.password); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
