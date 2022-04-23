FROM golang:1.18.1-bullseye AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o konfig main.go

FROM debian:bullseye-slim

WORKDIR /

RUN DEBIAN_FRONTEND=noninteractive apt-get update && apt-get -y upgrade && \
    apt-get install -y --no-install-recommends \
    ca-certificates && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

COPY --from=build /app/konfig /konfig
ENTRYPOINT ["/konfig"]
