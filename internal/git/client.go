package git

import (
	"net/url"

	"github.com/alexfalkowski/go-service/env"
	"github.com/alexfalkowski/go-service/telemetry/logger"
	"github.com/alexfalkowski/go-service/telemetry/metrics"
	"github.com/alexfalkowski/go-service/telemetry/tracer"
	"github.com/alexfalkowski/go-service/transport/http"
	"github.com/google/go-github/v69/github"
	"go.uber.org/fx"
)

// ConfigParams for git.
type ClientParams struct {
	fx.In

	Config    *http.Config
	Logger    *logger.Logger
	Endpoint  Endpoint
	Tracer    *tracer.Tracer
	Meter     *metrics.Meter
	UserAgent env.UserAgent
}

// NewClient for git.
func NewClient(params ClientParams) *github.Client {
	client, _ := http.NewClient(
		http.WithClientLogger(params.Logger), http.WithClientTracer(params.Tracer),
		http.WithClientMetrics(params.Meter), http.WithClientUserAgent(params.UserAgent),
		http.WithClientTimeout(params.Config.Timeout),
	)
	github := github.NewClient(client)
	endpoint := params.Endpoint

	if endpoint.IsSet() {
		u, _ := url.Parse(string(endpoint))
		github.BaseURL = u
	}

	return github
}
