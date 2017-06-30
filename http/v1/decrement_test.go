package v1

import (
	"testing"
)

func Test_client_Decrement(t *testing.T) {
	type args struct {
		cache string
		key   string
		init  *int64
		delta int64
	}
	tests := []struct {
		name       string
		c          Client
		args       args
		wantValue  int64
		wantNodeID string
		wantToken  string
		wantErr    bool
	}{
		{
			name: "Decrement atomic long",
			c:    NewClient([]string{"http://localhost:8080/ignite"}, "", ""),
			args: args{
				cache: "Person",
				key:   "sequence",
				init:  nil,
				delta: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, gotNodeID, gotToken, err := tt.c.Decrement(tt.args.cache, tt.args.key, tt.args.init, tt.args.delta)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.Decrement() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log("")
			t.Logf("http.v1.client.Decrement returned for '%s':", tt.name)
			t.Log("value =", gotValue)
			t.Log("affinityNodeId =", gotNodeID)
			t.Log("sessionToken =", gotToken)
		})
	}
}
