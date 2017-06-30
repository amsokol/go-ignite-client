package v1

import (
	"testing"

	core "github.com/amsokol/go-ignite-client/http"
)

func Test_client_GetCacheMetrics(t *testing.T) {
	type args struct {
		cache  string
		destID string
	}
	tests := []struct {
		name        string
		c           Client
		args        args
		wantMetrics core.CacheMetrics
		wantNodeID  string
		wantToken   string
		wantErr     bool
	}{
		{
			name: "Show metrics for Ignite cache",
			c:    NewClient([]string{"http://localhost:8080/ignite"}, "", ""),
			args: args{
				cache:  "Person",
				destID: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMetrics, gotNodeID, gotToken, err := tt.c.GetCacheMetrics(tt.args.cache, tt.args.destID)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.GetCacheMetrics() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log("")
			t.Logf("http.v1.client.GetCacheMetrics returned for '%s':", tt.name)
			t.Log("metrics =", gotMetrics)
			t.Log("affinityNodeId =", gotNodeID)
			t.Log("sessionToken =", gotToken)
		})
	}
}
