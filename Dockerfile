FROM golang:1.18.0-bullseye AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o konfig main.go

FROM debian:bullseye-slim

WORKDIR /
COPY --from=build /app/konfig /konfig
USER nonroot:nonroot
ENTRYPOINT ["/konfig"]
