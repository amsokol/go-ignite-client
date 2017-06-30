package internal

import (
	"sync"
	"testing"
)

func Test_pool_UpdateStatus(t *testing.T) {
	p := pool{data: make(map[string]bool), mx: &sync.Mutex{}}

	type args struct {
		server string
		alive  bool
	}
	tests := []struct {
		name string
		p    *pool
		args args
	}{
		{
			name: "Mark server is down",
			p:    &p,
			args: args{
				server: "http://localhost:8080/ignite",
				alive:  false,
			},
		},
		{
			name: "Mark server is alive",
			p:    &p,
			args: args{
				server: "http://localhost:8080/ignite",
				alive:  true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.UpdateStatus(tt.args.server, tt.args.alive)
		})
	}
}

func Test_pool_GetNextServer(t *testing.T) {
	p := pool{data: make(map[string]bool), mx: &sync.Mutex{}}
	servers := []string{"http://server1:8080/ignite", "http://server2:8080/ignite", "http://server3:8080/ignite"}

	p2 := pool{data: make(map[string]bool), mx: &sync.Mutex{}}
	p2.UpdateStatus("http://server1:8080/ignite", false)
	p2.UpdateStatus("http://server2:8080/ignite", false)
	p2.UpdateStatus("http://server3:8080/ignite", false)

	type args struct {
		server  string
		servers []string
	}
	tests := []struct {
		name    string
		p       *pool
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Empty server list",
			p:    &p,
			args: args{
				server:  "",
				servers: nil,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "First run",
			p:    &p,
			args: args{
				server:  "",
				servers: servers,
			},
			want: "http://server1:8080/ignite",
		},
		{
			name: "Second run",
			p:    &p,
			args: args{
				server:  "http://server1:8080/ignite",
				servers: servers,
			},
			want: "http://server2:8080/ignite",
		},
		{
			name: "Invalid last server",
			p:    &p,
			args: args{
				server:  "invalid value",
				servers: servers,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "First run but all servers are down",
			p:    &p2,
			args: args{
				server:  "",
				servers: servers,
			},
			want: "http://server1:8080/ignite",
		},
		{
			name: "Second run but all servers are down",
			p:    &p2,
			args: args{
				server:  "http://server1:8080/ignite",
				servers: servers,
			},
			want: "http://server2:8080/ignite",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.GetNextServer(tt.args.server, tt.args.servers)
			if (err != nil) != tt.wantErr {
				t.Errorf("pool.GetNextServer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("pool.GetNextServer() = %v, want %v", got, tt.want)
			}
		})
	}
}
