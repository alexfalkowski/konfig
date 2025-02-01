![Gopher](assets/gopher.png)
[![CircleCI](https://circleci.com/gh/alexfalkowski/konfig.svg?style=shield)](https://circleci.com/gh/alexfalkowski/konfig)
[![codecov](https://codecov.io/gh/alexfalkowski/konfig/graph/badge.svg?token=EZQ0Y4847T)](https://codecov.io/gh/alexfalkowski/konfig)
[![Go Report Card](https://goreportcard.com/badge/github.com/alexfalkowski/konfig)](https://goreportcard.com/report/github.com/alexfalkowski/konfig)
[![Go Reference](https://pkg.go.dev/badge/github.com/alexfalkowski/konfig.svg)](https://pkg.go.dev/github.com/alexfalkowski/konfig)
[![Stability: Active](https://masterminds.github.io/stability/active.svg)](https://masterminds.github.io/stability/active.html)

# Konfig

Konfig is a configuration system for application configuration.

## Background

Configuration is a very interesting topic. As we build more microservices we need to rethink how we get distributed systems to get their configuration. More info please read [External Configuration Store pattern](https://docs.microsoft.com/en-us/azure/architecture/patterns/external-configuration-store).

## Environment Variables

Well we have [environment variables](https://en.wikipedia.org/wiki/Environment_variable) so why do we need a whole service for this solved problem? That is a great question.

Here are some reasons:
- They are global state.
- The values cannot handle structures more complex than a string.
- They can't be versioned.
- They are hard to verify/validate for correctness.

## Configuration as Code

We want to standardize configuration and check it into version control. We are firm believers of using [GitOps](https://about.gitlab.com/topics/gitops/). Take a look at [Your configs suck? Try a real programming language](https://beepb00p.xyz/configs-suck.html). Some systems to have a look at:
- [The Dhall configuration language](https://dhall-lang.org/)
- [The Pkl configuration language](https://pkl-lang.org/index.html)

## Format

This system is geared around a very specific system that we use to build [services](https://github.com/alexfalkowski/konfig).

The kinds of this config that are supported are:
- [YAML](https://en.wikipedia.org/wiki/YAML)
- [TOML](https://en.wikipedia.org/wiki/TOML)

We recommend that you find a way to validate your configurations. We recommend looking at the following:
- [CUE](https://cuelang.org/)
- [yamllint](https://github.com/adrienverge/yamllint)

### Providers

The configuration can be augmented with values that might be sensitive and need to be retrieved at runtime.

#### Environment Variables

To retrieve an environment variables the value of the key in the config should be `env:VARIABLE`, ex: `env:GITHUB_URL`.

#### Vault

You can store values in [vault](https://learn.hashicorp.com/vault) for safe keeping.

##### Key

The key format is as follows:

```url
vault:/secret/data/key
```

An example:

```url
vault:/secret/data/transport/http/user_agent
```

##### Value

The value format is as follows:

```json
{"data": { "value": {} }}
```

An example:

```json
{"data": { "value": "Konfig-server/1.0 http/1.0" }}
```

##### Configuration

[Environment Variables](https://developer.hashicorp.com/vault/docs/commands#environment-variables)

#### SSM

You can store values in [ssm](https://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-parameter-store.html) for safe keeping.

##### Key

The key format is as follows:

```url
ssm:/secret/data/key
```

An example:

```url
ssm:/secret/data/transport/http/user_agent
```

##### Value

The value format is as follows:

```json
{"data": { "value": {} }}
```

An example:

```json
{"data": { "value": "Konfig-server/1.0 http/1.0" }}
```

##### Configuration

[Environment Variables](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-envvars.html)

## Source

This system allows you to store your configuration from various sources. Though we highly recommend that you follow configuration as code.

### Git

[Distributed version control](https://en.wikipedia.org/wiki/Distributed_version_control) is awesome and we believe should be used when managing configuration.

To configure we just need the have the following configuration:

```yaml
source:
  kind: git
  git:
    owner: the repo owner
    repository: the repo name
    token: path to token
```

We expect that the folders to have the following conventions:

```tree
application
└── environment
    ├── continent
    │   ├── country
    │   │   └── app.kind
    │   └── app.kind
    └── app.kind
```

The tag name should be `application/version` and kind is `yml`.

Some examples:
- [app-config](https://github.com/alexfalkowski/app-config)

### S3

[S3](https://aws.amazon.com/s3/) is another way to store your configurations.

To configure we just need the have the following configuration:

[Environment Variables](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-envvars.html)

```yaml
source:
  kind: s3
  s3:
    bucket: The bucket that contains all the configs.
```

We expect that the folders to have the following conventions:

```tree
application
└── version
    └── environment
        ├── continent
        │   ├── country
        │   │   └── app.kind
        │   └── app.kind
        └── app.kind
```

Some examples:

```url
s3://bucket/test/v1.5.0/production/server.kind
s3://bucket/test/v1.5.0/production/eu/server.kind
s3://bucket/test/v1.5.0/production/eu/de/server.kind
```

Kind is `yaml`, `toml`.

### Folder

This is mainly used for testing or if you want to quickly run it. If you have a secure way to mount these configs, then by all means go for it.

To configure we just need the have the following configuration:

```yaml
source:
  kind: folder
  folder:
    dir: .config (the folder where the configurations can be found)
```

We expect that the folders to have the following conventions:

```tree
application
└── version
    └── environment
        ├── continent
        │   ├── country
        │   │   └── app.kind
        │   └── app.kind
        └── app.kind
```

Kind is `yaml`, `toml`.

## Server

The server is defined by the following [proto contract](api/konfig/v1/service.proto). So each version of the service will have a new contract.

## Health

The system defines a way to monitor all of it's dependencies.

To configure we just need the have the following configuration:

```yaml
health:
  duration: 1s (how often to check)
  timeout: 1s (when we should timeout the check)
```

## Deployment

Since we are advocating building microservices, you would normally use a [container orchestration system](https://newrelic.com/blog/best-practices/container-orchestration-explained). Here is what we recommend when using this system:
- You could have a global config service or shard these config services per [bounded context](https://martinfowler.com/bliki/BoundedContext.html)
- The client should be used as an [init container](https://kubernetes.io/docs/concepts/workloads/pods/init-containers/).

## Other Systems

We love discovering systems that inspire us to make better systems. Below is a list of such systems:
- [Microconfig](https://microconfig.io/documentation.html#/)

## Development

If you would like to contribute, here is how you can get started.

### Structure

The project follows the structure in [golang-standards/project-layout](https://github.com/golang-standards/project-layout).

### Dependencies

Please make sure that you have the following installed:
- [Ruby](.ruby-version)
- Golang

### Setup

The get yourself setup, please run the following:

```sh
make setup
```

### Binaries

To make sure everything compiles for the app, please run the following:

```sh
make build-test
```

### Tests

To be able to test things locally you have to setup the environment.

#### Starting

Please run:

```sh
make start
```

#### Stopping

Please run:

```sh
make stop
```

#### Features

To run all the features, please run the following:

```sh
make features
```

### Changes

To see what has changed, please have a look at `CHANGELOG.md`
