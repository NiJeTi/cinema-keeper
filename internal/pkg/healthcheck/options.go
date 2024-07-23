package healthcheck

import (
	"fmt"
	"time"
)

type Option func(hc *Healthcheck)

func WithProbe(name string, probe Probe) Option {
	return func(hc *Healthcheck) {
		if _, ok := hc.probes[name]; ok {
			p := fmt.Sprintf("healthcheck probe '%s' already registered", name)
			panic(p)
		}

		hc.probes[name] = probe
	}
}

func WithTimeoutDegraded(timeout time.Duration) Option {
	return func(hc *Healthcheck) {
		if timeout <= 0 {
			panic("healthcheck timeout must be greater than zero")
		}

		hc.timeoutDegraded = timeout
	}
}

func WithTimeoutUnhealthy(timeout time.Duration) Option {
	return func(hc *Healthcheck) {
		if timeout <= 0 {
			panic("healthcheck timeout must be greater than zero")
		}

		hc.timeoutUnhealthy = timeout
	}
}
