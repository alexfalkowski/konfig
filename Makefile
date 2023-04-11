.PHONY: vendor

include bin/build/make/service.mak

# Build release binary.
build:
	go build -race -ldflags="-X 'github.com/alexfalkowski/konfig/cmd.Version=latest'" -mod vendor -o konfig main.go

# Build test binary.
build-test:
	go test -race -ldflags="-X 'github.com/alexfalkowski/konfig/cmd.Version=latest'" -mod vendor -c -tags features -covermode=atomic -o konfig -coverpkg=./... github.com/alexfalkowski/konfig

# Release to docker hub.
docker:
	bin/build/docker/push konfig
