package aws

import (
	aws "github.com/alexfalkowski/konfig/aws/endpoint"
	"github.com/alexfalkowski/konfig/aws/s3"
	"github.com/alexfalkowski/konfig/aws/ssm"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	fx.Provide(ssm.NewClient),
	fx.Provide(s3.NewClient),
	fx.Provide(aws.NewEndpoint),
)
