package http

import (
	"database/sql/driver"
	"testing"
)

func TestDriver_Open(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		a       *Driver
		args    args
		want    driver.Conn
		wantErr bool
	}{
		{
			name: "localhost",
			a:    nil,
			args: args{
				name: `{"servers" : ["http://localhost:8080/ignite"], "username" : "login", "password" : "password", "cache" : "Person"}`,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.a.Open(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Driver.Open() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
