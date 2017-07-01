package cache

import (
	core "github.com/amsokol/go-ignite-client/http"
	"github.com/amsokol/go-ignite-client/http/v1/exec"
)

// Management is interface to execute commands
type Management interface {
	GetOrCreateCache(e exec.Executer, cache string) (token string, err error)
	GetCacheMetrics(e exec.Executer, cache string, destID string) (metrics core.CacheMetrics, nodeID string, token string, err error)
}

// ManagementImpl is the implementation for Management interface
type ManagementImpl struct {
}
