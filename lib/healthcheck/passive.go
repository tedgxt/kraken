package healthcheck

import (
	"sync"
	"time"

	"code.uber.internal/infra/kraken/lib/hostlist"
	"code.uber.internal/infra/kraken/utils/stringset"

	"github.com/andres-erbsen/clock"
)

// Passive is a passive health check which tracks failed requests to hosts.
// Clients are responsible for marking failures from individual hosts, and
// Passive updates which hosts are unhealthy. It is recommended that clients
// only mark failures for network errors, not HTTP errors.
type Passive struct {
	sync.Mutex
	config    PassiveConfig
	clk       clock.Clock
	hosts     hostlist.List
	unhealthy map[string]time.Time
	failures  map[string][]time.Time
}

// NewPassive creates a new Passive wrapping hosts. See PassiveConfig for a
// description of how passive health checks are defined.
func NewPassive(config PassiveConfig, clk clock.Clock, hosts hostlist.List) *Passive {
	config.applyDefaults()
	return &Passive{
		config:    config,
		clk:       clk,
		hosts:     hosts,
		unhealthy: make(map[string]time.Time),
		failures:  make(map[string][]time.Time),
	}
}

// Resolve returns the latest healthy hosts. If all hosts are unhealthy, returns
// all hosts.
func (p *Passive) Resolve() stringset.Set {
	p.Lock()
	defer p.Unlock()

	healthy := p.hosts.Resolve()

	for addr, t := range p.unhealthy {
		if p.clk.Now().Sub(t) > p.config.FailTimeout {
			delete(p.unhealthy, addr)
		} else {
			healthy.Remove(addr)
		}
	}

	if len(healthy) == 0 {
		healthy = p.hosts.Resolve()
	}
	return healthy
}

// Failed marks a request to addr as failed.
func (p *Passive) Failed(addr string) {
	p.Lock()
	defer p.Unlock()

	now := p.clk.Now()

	failures := p.failures[addr]

	// Pop off the expired failures.
	for len(failures) > 0 {
		if now.Sub(failures[0]) > p.config.FailTimeout {
			failures = failures[1:]
			continue
		}
		break
	}

	// Add latest failure.
	failures = append(failures, now)

	if len(failures) >= p.config.Fails {
		p.unhealthy[addr] = now
	}
	p.failures[addr] = failures
}