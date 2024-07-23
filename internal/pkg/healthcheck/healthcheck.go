package healthcheck

import (
	"context"
	"log"
	"log/slog"
	"sync"
	"time"

	"github.com/valyala/fasthttp"
)

type Healthcheck struct {
	log              *slog.Logger
	probes           map[string]Probe
	timeoutDegraded  time.Duration
	timeoutUnhealthy time.Duration
}

func New(
	log *slog.Logger,
	opts ...Option,
) *Healthcheck {
	hc := &Healthcheck{
		log:              log,
		probes:           map[string]Probe{},
		timeoutDegraded:  1 * time.Second,
		timeoutUnhealthy: 10 * time.Second,
	}

	for _, opt := range opts {
		opt(hc)
	}

	return hc
}

func (hc *Healthcheck) Serve(address string) *fasthttp.Server {
	s := &fasthttp.Server{
		Handler: func(ctx *fasthttp.RequestCtx) {
			if string(ctx.Path()) != "/health" || !ctx.IsGet() {
				ctx.NotFound()
				return
			}

			status := hc.handle(ctx)

			if status == StatusHealthy {
				ctx.SetStatusCode(fasthttp.StatusOK)
			} else {
				ctx.SetStatusCode(fasthttp.StatusServiceUnavailable)
			}
			ctx.SetBodyString(status.String())
		},
	}

	go func() {
		err := s.ListenAndServe(address)
		if err != nil {
			log.Fatalln("healthcheck server error", "error", err)
		}
	}()

	return s
}

func (hc *Healthcheck) handle(ctx context.Context) Status {
	nProbes := len(hc.probes)

	wg := &sync.WaitGroup{}
	wg.Add(nProbes)

	statuses := make(chan Status, nProbes)

	for name, probe := range hc.probes {
		hc.probeCheck(
			hc.log.With("probe", name), ctx, wg, probe, statuses,
		)
	}

	wg.Wait()
	close(statuses)

	status := StatusHealthy
	for s := range statuses {
		if s <= status {
			continue
		}

		status = s
		if status == StatusUnhealthy {
			break
		}
	}

	return status
}

func (hc *Healthcheck) probeCheck(
	log *slog.Logger,
	ctx context.Context,
	wg *sync.WaitGroup,
	probe Probe,
	status chan Status,
) {
	timedCtx, cancel := context.WithTimeout(ctx, hc.timeoutUnhealthy)

	defer func() {
		if err := recover(); err != nil {
			log.ErrorContext(
				ctx,
				"probe panicked",
				"panic", err,
			)

			status <- StatusUnhealthy
		}

		cancel()
		wg.Done()
	}()

	probeTime := time.Now()
	err := probe.Check(timedCtx)
	probeDuration := time.Since(probeTime)

	if err != nil || probeDuration > hc.timeoutUnhealthy {
		log.ErrorContext(
			ctx,
			"failed to probe",
			"error", err,
			"duration", probeDuration.String(),
		)

		status <- StatusUnhealthy
		return
	}

	if probeDuration > hc.timeoutDegraded {
		log.WarnContext(
			ctx,
			"probe is degraded",
			"duration", probeDuration.String(),
		)

		status <- StatusDegraded
		return
	}

	status <- StatusHealthy
}
