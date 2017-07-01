package exec

import (
	"net/url"
)

// Executer is interface to execute commands
type Executer interface {
	Execute(v url.Values) (response []byte, nodeID string, token string, err error)
}

// ExecuterImpl is the implementation for Executer interface
type ExecuterImpl struct {
	Servers  []string
	Username string
	Password string
}
