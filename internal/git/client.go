package git

import (
	"net/url"

	"github.com/alexfalkowski/go-service/env"
	"github.com/alexfalkowski/go-service/transport/http"
	"github.com/google/go-github/v68/github"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// ConfigParams for git.
type ClientParams struct {
	fx.In

	Config    *http.Config
	Logger    *zap.Logger
	Endpoint  Endpoint
	Tracer    trace.Tracer
	Meter     metric.Meter
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
