FROM postgres:14.5 as db
WORKDIR /trackmywealth

FROM golang:1.19.2 AS development
WORKDIR /trackmywealth
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go install github.com/cespare/reflex@latest
EXPOSE 4000
CMD reflex -g '*.go' go run main.go --start-service

FROM golang:1.19.2 AS builder
ENV GOOS linux
ENV CGO_ENABLED 0
WORKDIR /trackmywealth
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o trackmywealth

FROM alpine:latest AS production
RUN apk add --no-cache ca-certificates
COPY --from=builder trackmywealth .
EXPOSE 4000
CMD ./trackmywealth