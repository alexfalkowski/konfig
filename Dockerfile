FROM golang:1.19.5-bullseye AS build

ARG version=latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build -ldflags="-X 'github.com/alexfalkowski/konfig/cmd.Version=${version}'" -a -o konfig main.go

FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=build /app/konfig /konfig
ENTRYPOINT ["/konfig"]
