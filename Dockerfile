FROM golang:1.23.5-bullseye AS build

ARG version=latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 go build -ldflags="-s -w -X 'github.com/alexfalkowski/konfig/cmd.Version=${version}'" -a -o konfig main.go

FROM gcr.io/distroless/static

WORKDIR /

COPY --from=build /app/konfig /konfig
ENTRYPOINT ["/konfig"]
