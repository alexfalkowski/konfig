include bin/build/make/service.mak
include bin/build/make/git.mak

# Setup secret.
setup-secret:
	echo ${GITHUB_TOKEN} > test/secrets/gh

# Run all gRPC features.
features-grpc:
	$(MAKE) feature=features tags=@grpc features

# Run all HTTP features.
features-http:
	$(MAKE) feature=features tags=@http features
