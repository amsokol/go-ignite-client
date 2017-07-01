package cache

import (
	"testing"

	core "github.com/amsokol/go-ignite-client/http"
	"github.com/amsokol/go-ignite-client/http/v1/exec"
)

func TestManagementImpl_GetCacheMetrics(t *testing.T) {
	t.Log("")
	t.Log("Preparing test data for 'TestManagementImpl_GetCacheMetrics'...")

	e := exec.ExecuterImpl{Servers: []string{"http://localhost:8080/ignite"}, Username: "", Password: ""}
	c := ManagementImpl{}

	_, err := c.GetOrCreateCache(&e, "Cache4TestGetOrCreateCache")
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Done")

	type args struct {
		e      exec.Executer
		cache  string
		destID string
	}
	tests := []struct {
		name        string
		c           *ManagementImpl
		args        args
		wantMetrics core.CacheMetrics
		wantNodeID  string
		wantToken   string
		wantErr     bool
	}{
		{
			name: "Get metrics for cache Cache4TestGetOrCreateCache",
			c:    &c,
			args: args{
				e:     &e,
				cache: "Cache4TestGetOrCreateCache",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMetrics, gotNodeID, gotToken, err := tt.c.GetCacheMetrics(tt.args.e, tt.args.cache, tt.args.destID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ManagementImpl.GetCacheMetrics() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log("")
			t.Logf("http.v1.cache.GetCacheMetrics returned for '%s':", tt.name)
			t.Log("metrics =", gotMetrics)
			t.Log("affinityNodeId =", gotNodeID)
			t.Log("sessionToken =", gotToken)
		})
	}
}
