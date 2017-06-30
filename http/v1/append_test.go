package v1

import "testing"

func Test_client_Append(t *testing.T) {
	type args struct {
		cache  string
		key    string
		val    string
		destID string
	}
	tests := []struct {
		name       string
		c          Client
		args       args
		wantOk     bool
		wantNodeID string
		wantToken  string
		wantErr    bool
	}{
		{
			name: "1",
			c:    NewClient([]string{"http://localhost:8080/ignite"}, "", ""),
			args: args{
				cache: "Cache1",
				key:   "key",
				val:   "new value",
			},
			wantOk:  false,
			wantErr: true,
		},
		// TODO: add more tests
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOk, _, _, err := tt.c.Append(tt.args.cache, tt.args.key, tt.args.val, tt.args.destID)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.Append() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOk != tt.wantOk {
				t.Errorf("client.Append() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
