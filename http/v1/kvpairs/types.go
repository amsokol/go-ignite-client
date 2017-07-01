package kvpairs

import (
	"github.com/amsokol/go-ignite-client/http/v1/exec"
)

// KeyValuePairs is interface to execute commands
type KeyValuePairs interface {
	Add(e exec.Executer, cache string, key string, val string, destID string) (ok bool, nodeID string, token string, err error)
	Append(e exec.Executer, cache string, key string, val string, destID string) (ok bool, nodeID string, token string, err error)
	CompareAndSwap(e exec.Executer, cache string, key string, val string, val2 string, destID string) (ok bool, nodeID string, token string, err error)
	Prepend(e exec.Executer, cache string, key string, val string, destID string) (ok bool, nodeID string, token string, err error)
	Replace(e exec.Executer, cache string, key string, val string, val2 string, destID string) (ok bool, nodeID string, token string, err error)
}

// KeyValuePairsImpl is the implementation for KeyValuePairs interface
type KeyValuePairsImpl struct {
}
