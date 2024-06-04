package ssm

import (
	"context"
	"net/http"
	"net/url"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/ssm"
	endpoints "github.com/aws/smithy-go/endpoints"
)

type resolver struct {
	ssm.EndpointResolverV2
}

func (r *resolver) ResolveEndpoint(ctx context.Context, params ssm.EndpointParameters) (endpoints.Endpoint, error) {
	u := os.Getenv("AWS_URL")
	if u != "" {
		p, err := url.Parse(u)
		if err != nil {
			return endpoints.Endpoint{}, err
		}

		return endpoints.Endpoint{URI: *p, Headers: http.Header{}}, nil
	}

	return r.EndpointResolverV2.ResolveEndpoint(ctx, params)
}
