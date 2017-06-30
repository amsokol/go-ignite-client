package v1

import "testing"

func Test_client_Replace(t *testing.T) {
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOk, _, _, err := tt.c.Replace(tt.args.cache, tt.args.key, tt.args.val, tt.args.destID)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.Replace() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOk != tt.wantOk {
				t.Errorf("client.Replace() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
