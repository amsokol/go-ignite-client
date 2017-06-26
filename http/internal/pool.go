package internal

import (
	"io"
	// "log"
	"math"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type node struct {
	lastupdate time.Time
	alive      bool
}

func (n *node) isQuarantineExpired(quarantine float64) bool {
	return math.Abs(time.Now().Sub(n.lastupdate).Minutes()) > quarantine
	//	d := time.Now().Sub(n.lastupdate).Minutes()
	//	log.Println("Mins from last =", d, " quarantine =", quarantine)
	//	return d > quarantine
}

type pool struct {
	data map[string]node
}

// GlobalPool is global pool
var GlobalPool = pool{data: make(map[string]node)}

func (p *pool) UpdateStatus(server string, alive bool) {
	p.data[server] = node{lastupdate: time.Now(), alive: alive}
	//	log.Println("Status of server", server, "set alive =", alive)
}

func (p *pool) GetNextServer(server string, servers []string, quarantine float64) (string, error) {
	if len(servers) == 0 {
		return "", errors.New("Your server list is empty")
	}

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

	// try to find node with no status, alive nodes or where quarantine is expired:
	for i := 0; i < count; i++ {
		index := (i + start) % size
		s := servers[index]
		node, found := p.data[s]
		if !found {
			return s, nil
		}
		if node.alive {
			return s, nil
		}
		if node.isQuarantineExpired(quarantine) {
			return s, nil
		}
	}
	return "", io.EOF
}
