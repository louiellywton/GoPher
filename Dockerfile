# Multi-stage Dockerfile for hello-gopher CLI
# This creates a minimal container image for the hello-gopher CLI tool

# Build stage - use official Go image for building
FROM golang:1.21-alpine AS builder

# Set working directory
WORKDIR /app

# Install git (needed for Go modules)
RUN apk add --no-cache git

# Copy go mod files first for better caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the binary with optimizations
# - CGO_ENABLED=0 for static binary
# - -ldflags for smaller binary size and version info
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-s -w -X github.com/louiellywton/go-portfolio/01-hello-gopher/cmd/hello-gopher/cmd.version=docker" \
    -o hello-gopher \
    ./cmd/hello-gopher

# Runtime stage - use minimal scratch image
FROM scratch

# Copy the binary from builder stage
COPY --from=builder /app/hello-gopher /hello-gopher

# Set the binary as entrypoint
ENTRYPOINT ["/hello-gopher"]

# Default command shows help
CMD ["--help"]