include bin/build/make/grpc.mak
include bin/build/make/git.mak

# Setup secret.
setup-secret:
	echo ${GITHUB_TOKEN} > test/secrets/gh

# Run all config features.
features-config:
	$(MAKE) feature=features tags=@config features

# Run all secrets features.
features-secrets:
	$(MAKE) feature=features tags=@secrets features
