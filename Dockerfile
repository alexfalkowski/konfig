FROM golang:1.22.1-bullseye AS build

ARG version=latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build -ldflags="-s -w -X 'github.com/alexfalkowski/konfig/cmd.Version=${version}'" -a -o konfig main.go

FROM gcr.io/distroless/base-debian12

WORKDIR /

COPY --from=build /app/konfig /konfig
ENTRYPOINT ["/konfig"]
