package v10

import (
	"log"
	"testing"

	"github.com/amsokol/go-ignite-client/http/types"
	"github.com/amsokol/go-ignite-client/http/v1/client"
)

func TestCacheMetrics(t *testing.T) {
	type args struct {
		c         client.Client
		cacheName string
		destID    string
	}
	tests := []struct {
		name               string
		args               args
		wantMetrics        types.CacheMetrics
		wantAffinityNodeID string
		wantSessionToken   types.SessionToken
		wantErr            bool
	}{
		{
			name: "Show metrics for Ignite cache",
			args: args{
				c:         client.Open([]string{"http://localhost:8080/ignite"}, "", ""),
				cacheName: "Person",
				destID:    "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMetrics, gotAffinityNodeID, gotSessionToken, err := CacheMetrics(tt.args.c, tt.args.cacheName, tt.args.destID)
			if (err != nil) != tt.wantErr {
				t.Errorf("CacheMetrics() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			log.Println("")
			log.Println("http.v1.v10.CacheMetrics returned:")
			log.Println("metrics=", gotMetrics)
			log.Println("affinityNodeId=", gotAffinityNodeID)
			log.Println("sessionToken=", gotSessionToken)
		})
	}
}
