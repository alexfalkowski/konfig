package aws

import (
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
)

// Endpoint for AWS.
func EndpointResolver() aws.EndpointResolverWithOptionsFunc {
	resolver := aws.EndpointResolverWithOptionsFunc(func(_, region string, _ ...any) (aws.Endpoint, error) {
		url := os.Getenv("AWS_URL")
		if url != "" {
			return aws.Endpoint{PartitionID: "aws", URL: url, SigningRegion: region}, nil
		}

		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})

	return resolver
}
