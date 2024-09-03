package hcServer

import (
	"log"

	"github.com/nijeti/healthcheck"
	"github.com/valyala/fasthttp"
)

func Serve(address string, hc *healthcheck.Healthcheck) *fasthttp.Server {
	s := &fasthttp.Server{
		GetOnly: true,
		Handler: func(ctx *fasthttp.RequestCtx) {
			if string(ctx.Path()) != "/health" {
				ctx.NotFound()
				return
			}

			status := hc.Handle(ctx)

			if status == healthcheck.StatusHealthy {
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
