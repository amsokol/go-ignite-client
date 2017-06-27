package internal

import (
	"strings"
	"sync"

	"github.com/pkg/errors"
)

type pool struct {
	data map[string]bool
	mx   *sync.Mutex
}

// GlobalPool is global pool
var GlobalPool = pool{data: make(map[string]bool), mx: &sync.Mutex{}}

func (p *pool) UpdateStatus(server string, alive bool) {
	p.mx.Lock()
	p.data[server] = alive
	p.mx.Unlock()
}

func (p *pool) GetNextServer(server string, servers []string) (string, error) {
	if len(servers) == 0 {
		return "", errors.New("Your server list is empty")
	}

	// find next server
	start := -1
	size := len(servers)
	count := size
	if len(server) > 0 {
		for i, s := range servers {
			if server == s {
				start = i + 1
				break
			}
		}
		if start == -1 {
			return "", errors.New(strings.Join([]string{"Server not found in your server list: ", server}, ""))
		}
		count--
	} else {
		start = 0
	}

	// try to find node with no status or alive nodes:
	for i := 0; i < count; i++ {
		index := (i + start) % size
		s := servers[index]
		p.mx.Lock()
		alive, found := p.data[s]
		p.mx.Unlock()
		if !found || alive {
			return s, nil
		}
	}

	// return next server
	return servers[start%size], nil
}
