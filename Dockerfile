FROM golang:1.18.0-bullseye AS build

WORKDIR /app

COPY Makefile ./
COPY go.mod ./
COPY go.sum ./
RUN make go-dep

COPY . ./
RUN make build

FROM debian:bullseye-slim

WORKDIR /
COPY --from=build /app/konfig /konfig
USER nonroot:nonroot
ENTRYPOINT ["/konfig"]
