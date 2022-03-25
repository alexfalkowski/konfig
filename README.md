[![CircleCI](https://circleci.com/gh/alexfalkowski/konfig.svg?style=svg)](https://circleci.com/gh/alexfalkowski/konfig)
[![Coverage Status](https://coveralls.io/repos/github/alexfalkowski/konfig/badge.svg?branch=master)](https://coveralls.io/github/alexfalkowski/konfig?branch=master)

# Konfig

Konfig is a configuration system for application configuration.

## Structure

The project follows the structure in [golang-standards/project-layout](https://github.com/golang-standards/project-layout)

## Dependencies

Please make sure that you have the following installed:
- Ruby (look at the .ruby-version)
- Golang

## Help

To see what can be run, please run the following:
```sh
make help
```

## Setup

The get yourself setup, please run the following:
```sh
make setup
```

## Binaries

To make sure everything compiles for the app, please run the following:
```sh
make build-test
```

## Tests

To be able to test things locally you have to setup the environment.

### Starting

Please run:

```sh
make start
```

### Stopping

Please run:

```sh
make stop
```

### Features

To run all the features, please run the following:
```sh
make features
```

## Changes

To see what has changed, please have a look at `CHANGELOG.md`
