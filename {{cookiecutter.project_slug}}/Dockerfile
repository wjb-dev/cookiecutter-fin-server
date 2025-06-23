# ──────────────────── Stage 1: Builder ────────────────────────────────────────
FROM golang:1.24.4-alpine AS builder

# Install protoc and Git
RUN apk add --no-cache git protobuf protoc

# Install the Go protoc plugins into $GOPATH/bin
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Ensure the plugins are on PATH
ENV PATH="${PATH}:$(go env GOPATH)/bin"


WORKDIR /app
# Copy all source code
COPY . .
# Copy go.mod and go.sum first to leverage Docker cache
COPY go.mod go.sum ./
RUN go get gopkg.in/yaml.v3@v3.0.1

RUN go mod download

# Generate Go code from your proto
RUN protoc \
-I pkg/proto/v1 \
--go_out=paths=source_relative:pkg/proto/v1 \
--go-grpc_out=paths=source_relative:pkg/proto/v1 \
pkg/proto/v1/service.proto

RUN go mod tidy

# Build static Go binary
RUN CGO_ENABLED=0 go build -o /app/{{ cookiecutter.project_slug }} cmd/{{ cookiecutter.project_slug }}/main.go


# ──────────────────── Stage 2: Runtime ────────────────────────────────────────
FROM alpine:latest

# Only CA certs (for TLS) & your binary
RUN apk add --no-cache ca-certificates

WORKDIR /app
COPY --from=builder /app/{{ cookiecutter.project_slug }} .
# Copy configs directory
COPY --from=builder /app/configs /app/configs

EXPOSE 50051
ENTRYPOINT ["./{{ cookiecutter.project_slug }}"]
