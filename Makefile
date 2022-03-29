.PHONY: vendor tools

help: ## Display this help
	@ echo "Please use \`make <target>' where <target> is one of:"
	@ echo
	@ grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "    \033[36m%-10s\033[0m - %s\n", $$1, $$2}'
	@ echo

tools: ## Setup all the tools
	tools/deps
	tools/googleapis

ruby-setup: ## Setup ruby
	make -C test setup

ruby-dep: ## Setup ruby deps
	make -C test dep

setup: tools go-dep ruby-setup ruby-dep ## Setup everything

download:
	go mod download

tidy:
	go mod tidy

vendor:
	go mod vendor

go-dep: download tidy vendor ## Setup go deps

dep: go-dep ruby-dep ## Setup all deps

build: ## Build release binary
	go build -mod vendor -o konfig main.go

build-test: ## Build test binary
	go test -mod vendor -c -tags features -covermode=count -o konfig -coverpkg=./... github.com/alexfalkowski/konfig

go-lint: ## Lint all the go code
	golangci-lint run --build-tags features --timeout 5m

go-fix-lint: ## Fix the lint issues in the go code (if possible)
	golangci-lint run --build-tags features --timeout 5m --fix

ruby-lint: ## Lint all the ruby code
	make -C test lint

ruby-fix-lint: ## Fix the lint issues in the ruby code (if possible)
	make -C test fix-lint

lint: go-lint ruby-lint ## Lint all the code

fix-lint: go-fix-lint ruby-fix-lint ## Fix the lint issues in the code (if possible)

features: build-test ## Run all the features
	make -C test features

sanitize-coverage:
	./tools/coverage

html-coverage: sanitize-coverage ## Get the HTML coverage for go
	go tool cover -html test/reports/final.cov

func-coverage: sanitize-coverage ## Get the func coverage for go
	go tool cover -func test/reports/final.cov

goveralls: sanitize-coverage ## Send coveralls data
	goveralls -coverprofile=test/reports/final.cov -service=circle-ci -repotoken=1r7TP3L2xhnSiOOutstLIB306z67K120W

ruby-generate-proto: ## Generate proto for ruby
	make -C test generate-proto

go-generate-proto: ## Generate proto for go
	tools/protoc

generate-proto: go-generate-proto ruby-generate-proto ## Generate proto

ruby-outdated: ## Check outdated ruby deps
	make -C test outdated

go-outdated: ## Check outdated go deps
	go list -u -m -mod=mod -json all | go-mod-outdated -update -direct

outdated: go-outdated ruby-outdated  ## Check outdated deps

go-get: ## Get go dep
	go get $(module)

go-update-dep: go-get tidy vendor ## Update go dep

ruby-update-dep: ## Update ruby dep
	make -C test gem=$(gem) update-dep

start: ## Start the environment
	tools/env start

stop: ## Stop the environment
	tools/env stop
