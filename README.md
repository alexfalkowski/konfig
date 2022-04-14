[![CircleCI](https://circleci.com/gh/alexfalkowski/konfig.svg?style=svg)](https://circleci.com/gh/alexfalkowski/konfig)
[![Coverage Status](https://coveralls.io/repos/github/alexfalkowski/konfig/badge.svg?branch=master)](https://coveralls.io/github/alexfalkowski/konfig?branch=master)

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

## Format

This system is geared around a very specific system that we use to build [services](https://github.com/alexfalkowski/go-service). The content type of this config is [YAML](https://en.wikipedia.org/wiki/YAML).

We recommend that you find a way to validate your configurations. We recommend looking at the following:
- [CUE](https://cuelang.org/)
- [yamllint](https://github.com/adrienverge/yamllint)

### Providers

The configuration can be augmented with values that might be sensitive and need to be retrieved at runtime.

#### Environment Variables

To retrieve an environment variables the value of the key in the config should be `env:VARIABLE`, ex: `env:GITHUB_URL`.

#### Vault

You can also store values in [vault](https://learn.hashicorp.com/vault) for safe keeping. To retrieve the value of the key in the config should be `vault:secret/data/key`, ex: `vault:secret/data/transport/http/user_agent`.

## Server

The server is defined by the following [proto contract](api/konfig/v1/service.proto). So each version of the service will have a new contract.

To configure we just need the have the following configuration:

```yaml
server:
  v1:
    source:
      type: git or folder (see below)
```

### Source

This system allows you to store your configuration from various sources. Though we highly recommend that you follow configuration as code.

#### Git

[Distributed version control](https://en.wikipedia.org/wiki/Distributed_version_control) is awesome and we believe should be used when managing configuration.

To configure we just need the have the following configuration:

```yaml
source:
    type: git
    git:
        url: https://github.com/alexfalkowski/app-config (the configuration repo)
        dir: tmp/app-config (where to clone the repo to)
        token: a GitHub token or can be set in KONFIG_GIT_TOKEN env variable
```

We expect the repo to have the following conventions:
- Each application name is at the root of the repository.
- Under the application we have the environments (staging, production, etc).
- Under each environment we have the configuration that follows `command.config.yml`. Where command should follow the commands that your service has. Like server, worker, etc.
- Versions are tracked by having the name of the service and the version. So a tag would would look like `test/v1.5.0`.

Take a look at [app-config)](https://github.com/alexfalkowski/app-config) as an example.

#### Folder

This is mainly used for testing or if you want to quickly run it. If you have a secure way to mount these configs, then by all means go for it.

To configure we just need the have the following configuration:

```yaml
source:
    type: folder
    folder:
        dir: .config (the folder where the configurations can be found)
```

We expect that the folders to have the following conventions:
- Each application name is at the root of the folder.
- Each version is under the application and is in the format of `v1.5.0`
- Under the version we have the environments (staging, production, etc).
- Under each environment we have the configuration that follows `command.config.yml`. Where command should follow the commands that your service has. Like server, worker, etc.

## Client

The client is used to get the config that is defined in the config. These values reflect how the config is stored in the above sources.

To configure we just need the have the following configuration:

```yaml
client:
  host: localhost:9090
  application: test
  version: v1.5.0
  environment: staging
  command: server
```

The client writes the config to the location specified by `APP_CONFIG_FILE` environment variable.

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
- Ruby (look at the .ruby-version)
- Golang

### Help

To see what can be run, please run the following:

```sh
make help
```

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
