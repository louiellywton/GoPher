# Build Guide - Hello-Gopher CLI

This document provides comprehensive instructions for building, testing, and releasing the Hello-Gopher CLI tool. It serves as both a development guide and a portfolio demonstration of professional Go project management.

## 📋 Table of Contents

- [Prerequisites](#prerequisites)
- [Quick Start](#quick-start)
- [Development Environment Setup](#development-environment-setup)
- [Building the Application](#building-the-application)
- [Testing Strategy](#testing-strategy)
- [Code Quality & Linting](#code-quality--linting)
- [Cross-Platform Building](#cross-platform-building)
- [Release Process](#release-process)
- [CI/CD Pipeline](#cicd-pipeline)
- [Troubleshooting](#troubleshooting)
- [Performance Optimization](#performance-optimization)

## 🔧 Prerequisites

### Required Tools

- **Go**: Version 1.21 or later
  ```bash
  # Check your Go version
  go version
  
  # Install Go if needed (example for Linux/macOS)
  # Visit https://golang.org/dl/ for official installers
  ```

- **Git**: For version control
  ```bash
  git --version
  ```

### Optional Tools (Recommended)

- **golangci-lint**: For comprehensive linting
  ```bash
  # Install golangci-lint
  go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
  ```

- **gosec**: For security scanning
  ```bash
  go install github.com/securecodewarrior/gosec/v2/cmd/gosec@latest
  ```

- **govulncheck**: For vulnerability scanning
  ```bash
  go install golang.org/x/vuln/cmd/govulncheck@latest
  ```

- **goreleaser**: For release automation (optional for development)
  ```bash
  # Install via Homebrew (macOS/Linux)
  brew install goreleaser
  
  # Or via Go
  go install github.com/goreleaser/goreleaser@latest
  ```

## 🚀 Quick Start

Get up and running in under 2 minutes:

```bash
# 1. Clone the repository
git clone https://github.com/louiellywton/go-portfolio.git
cd go-portfolio/01-hello-gopher

# 2. Download dependencies
go mod download

# 3. Build and run
go build -o hello-gopher ./cmd/hello-gopher
./hello-gopher greet --name "Builder"

# 4. Run tests
go test ./... -race -cover
```

## 🛠️ Development Environment Setup

### 1. Clone and Navigate

```bash
git clone https://github.com/louiellywton/go-portfolio.git
cd go-portfolio/01-hello-gopher
```

### 2. Verify Project Structure

```bash
tree -I 'vendor|.git'
```

Expected structure:
```
01-hello-gopher/
├── cmd/
│   └── hello-gopher/
│       ├── main.go
│       └── cmd/
├── pkg/
│   └── greeting/
├── scripts/
├── .github/
├── .goreleaser.yaml
├── go.mod
├── go.sum
└── README.md
```

### 3. Dependency Management

```bash
# Download all dependencies
go mod download

# Verify dependencies
go mod verify

# Clean up unused dependencies
go mod tidy

# View dependency graph
go mod graph
```

### 4. Environment Variables

Set up optional environment variables for development:

```bash
# Enable Go modules (default in Go 1.16+)
export GO111MODULE=on

# Enable detailed error messages
export GOTRACEBACK=all

# Enable race detector (for testing)
export GORACE="log_path=./race_log"
```

## 🔨 Building the Application

### Development Build

```bash
# Simple build
go build -o hello-gopher ./cmd/hello-gopher

# Build with race detector (for testing)
go build -race -o hello-gopher-race ./cmd/hello-gopher

# Build with debug information
go build -gcflags="all=-N -l" -o hello-gopher-debug ./cmd/hello-gopher
```

### Production Build

```bash
# Optimized build with version information
go build \
  -ldflags="-s -w -X github.com/louiellywton/go-portfolio/01-hello-gopher/cmd/hello-gopher/cmd.version=v1.0.0 -X github.com/louiellywton/go-portfolio/01-hello-gopher/cmd/hello-gopher/cmd.buildDate=$(date -u +%Y-%m-%dT%H:%M:%SZ) -X github.com/louiellywton/go-portfolio/01-hello-gopher/cmd/hello-gopher/cmd.gitCommit=$(git rev-parse HEAD)" \
  -o hello-gopher \
  ./cmd/hello-gopher

# Verify the build
./hello-gopher --version
```

### Build Flags Explained

- `-ldflags="-s -w"`: Strip debug info and symbol table (reduces binary size)
- `-X package.variable=value`: Set string variable at build time
- `-race`: Enable race detector (development only)
- `-gcflags="all=-N -l"`: Disable optimizations for debugging

## 🧪 Testing Strategy

### Unit Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test ./... -cover

# Run tests with race detection
go test ./... -race

# Verbose test output
go test -v ./...

# Run specific test
go test -run TestGreet ./pkg/greeting

# Run tests in specific package
go test ./pkg/greeting/...
```

### Coverage Analysis

```bash
# Generate coverage profile
go test ./... -coverprofile=coverage.out

# View coverage in terminal
go tool cover -func=coverage.out

# Generate HTML coverage report
go tool cover -html=coverage.out -o coverage.html

# Open coverage report in browser
open coverage.html  # macOS
xdg-open coverage.html  # Linux
start coverage.html  # Windows
```

### Benchmark Tests

```bash
# Run all benchmarks
go test -bench=. ./...

# Run benchmarks with memory allocation stats
go test -bench=. -benchmem ./...

# Run specific benchmark
go test -bench=BenchmarkGreet ./pkg/greeting

# Run benchmarks multiple times for accuracy
go test -bench=. -count=5 ./...
```

### Test Categories

The project includes several types of tests:

1. **Unit Tests**: Test individual functions in isolation
2. **Integration Tests**: Test command execution end-to-end
3. **Table-Driven Tests**: Test multiple scenarios efficiently
4. **Mock Tests**: Test with mocked dependencies
5. **Example Tests**: Executable documentation
6. **Benchmark Tests**: Performance validation

### Example Test Commands

```bash
# Test with timeout
go test -timeout 30s ./...

# Test with short flag (skip long-running tests)
go test -short ./...

# Test with custom flags
go test -args -custom-flag=value ./...

# Parallel test execution
go test -parallel 4 ./...
```

## 🔍 Code Quality & Linting

### golangci-lint

```bash
# Run all linters
golangci-lint run

# Run with specific linters
golangci-lint run --enable=gosec,gocritic

# Run with configuration file
golangci-lint run --config .golangci.yml

# Fix auto-fixable issues
golangci-lint run --fix
```

### Security Scanning

```bash
# Run gosec security scanner
gosec ./...

# Run with JSON output
gosec -fmt json -out gosec-report.json ./...

# Check for vulnerabilities
govulncheck ./...
```

### Code Formatting

```bash
# Format all Go files
go fmt ./...

# Check if files are formatted
gofmt -l .

# Format with imports organization
goimports -w .
```

### Static Analysis

```bash
# Run go vet
go vet ./...

# Check for inefficient assignments
ineffassign ./...

# Check for misspellings
misspell .

# Check for unused code
deadcode ./...
```

## 🌍 Cross-Platform Building

### Manual Cross-Platform Builds

```bash
# Linux AMD64
GOOS=linux GOARCH=amd64 go build -o hello-gopher-linux-amd64 ./cmd/hello-gopher

# Linux ARM64
GOOS=linux GOARCH=arm64 go build -o hello-gopher-linux-arm64 ./cmd/hello-gopher

# macOS AMD64
GOOS=darwin GOARCH=amd64 go build -o hello-gopher-darwin-amd64 ./cmd/hello-gopher

# macOS ARM64 (Apple Silicon)
GOOS=darwin GOARCH=arm64 go build -o hello-gopher-darwin-arm64 ./cmd/hello-gopher

# Windows AMD64
GOOS=windows GOARCH=amd64 go build -o hello-gopher-windows-amd64.exe ./cmd/hello-gopher

# Windows ARM64
GOOS=windows GOARCH=arm64 go build -o hello-gopher-windows-arm64.exe ./cmd/hello-gopher
```

### Build Script

Create a build script for convenience:

```bash
#!/bin/bash
# build-all.sh

set -e

VERSION=${1:-dev}
BUILD_DATE=$(date -u +%Y-%m-%dT%H:%M:%SZ)
GIT_COMMIT=$(git rev-parse HEAD)

LDFLAGS="-s -w -X github.com/louiellywton/go-portfolio/01-hello-gopher/cmd/hello-gopher/cmd.version=${VERSION} -X github.com/louiellywton/go-portfolio/01-hello-gopher/cmd/hello-gopher/cmd.buildDate=${BUILD_DATE} -X github.com/louiellywton/go-portfolio/01-hello-gopher/cmd/hello-gopher/cmd.gitCommit=${GIT_COMMIT}"

echo "Building hello-gopher ${VERSION}..."

# Create build directory
mkdir -p build

# Build for all platforms
GOOS=linux GOARCH=amd64 go build -ldflags="${LDFLAGS}" -o build/hello-gopher-linux-amd64 ./cmd/hello-gopher
GOOS=linux GOARCH=arm64 go build -ldflags="${LDFLAGS}" -o build/hello-gopher-linux-arm64 ./cmd/hello-gopher
GOOS=darwin GOARCH=amd64 go build -ldflags="${LDFLAGS}" -o build/hello-gopher-darwin-amd64 ./cmd/hello-gopher
GOOS=darwin GOARCH=arm64 go build -ldflags="${LDFLAGS}" -o build/hello-gopher-darwin-arm64 ./cmd/hello-gopher
GOOS=windows GOARCH=amd64 go build -ldflags="${LDFLAGS}" -o build/hello-gopher-windows-amd64.exe ./cmd/hello-gopher
GOOS=windows GOARCH=arm64 go build -ldflags="${LDFLAGS}" -o build/hello-gopher-windows-arm64.exe ./cmd/hello-gopher

echo "Build complete! Binaries are in the build/ directory."
ls -la build/
```

Make it executable and run:

```bash
chmod +x build-all.sh
./build-all.sh v1.0.0
```

## 🚢 Release Process

### Using Goreleaser

The project uses Goreleaser for automated releases:

```bash
# Test release configuration
goreleaser check

# Build snapshot (without releasing)
goreleaser build --snapshot --clean

# Create a test release
goreleaser release --snapshot --clean

# Create actual release (requires git tag)
git tag v1.0.0
git push origin v1.0.0
goreleaser release --clean
```

### Manual Release Process

1. **Update Version**:
   ```bash
   # Update version in relevant files
   git add .
   git commit -m "Bump version to v1.0.0"
   ```

2. **Create Git Tag**:
   ```bash
   git tag -a v1.0.0 -m "Release v1.0.0"
   git push origin v1.0.0
   ```

3. **Build Release Binaries**:
   ```bash
   ./build-all.sh v1.0.0
   ```

4. **Create Release Archives**:
   ```bash
   # Create archives for distribution
   cd build
   tar -czf hello-gopher-v1.0.0-linux-amd64.tar.gz hello-gopher-linux-amd64
   tar -czf hello-gopher-v1.0.0-darwin-amd64.tar.gz hello-gopher-darwin-amd64
   zip hello-gopher-v1.0.0-windows-amd64.zip hello-gopher-windows-amd64.exe
   ```

5. **Generate Checksums**:
   ```bash
   sha256sum *.tar.gz *.zip > checksums.txt
   ```

### Release Checklist

- [ ] All tests pass
- [ ] Code coverage meets threshold (80%+)
- [ ] Linting passes
- [ ] Security scan passes
- [ ] Documentation updated
- [ ] Version bumped
- [ ] Git tag created
- [ ] Release notes prepared
- [ ] Binaries built and tested
- [ ] Homebrew tap updated (automated)

## 🔄 CI/CD Pipeline

### GitHub Actions Workflow

The project uses GitHub Actions for CI/CD with the following jobs:

1. **Test**: Multi-platform testing with Go version matrix
2. **Lint**: Code quality checks with golangci-lint
3. **Security**: Security scanning with gosec and govulncheck
4. **Build**: Cross-platform binary building
5. **Coverage**: Test coverage reporting
6. **Benchmark**: Performance benchmarking
7. **Release**: Automated releases with Goreleaser

### Local CI Simulation

```bash
# Simulate the CI pipeline locally
./scripts/test-ci.sh  # or test-ci.ps1 on Windows
```

### Pipeline Configuration

The CI pipeline is configured in `.github/workflows/ci.yml` and includes:

- **Multi-OS Testing**: Ubuntu, macOS, Windows
- **Go Version Matrix**: 1.21, 1.22
- **Race Detection**: Concurrent safety testing
- **Coverage Reporting**: Codecov integration
- **Artifact Upload**: Build artifacts and reports
- **Release Automation**: Goreleaser integration

## 🐛 Troubleshooting

### Common Build Issues

#### Go Module Issues

```bash
# Problem: Module not found
go mod download
go mod verify

# Problem: Dependency conflicts
go mod tidy
go clean -modcache
```

#### Build Failures

```bash
# Problem: CGO errors
export CGO_ENABLED=0
go build ./cmd/hello-gopher

# Problem: Missing dependencies
go mod download
go mod verify
```

#### Test Failures

```bash
# Problem: Race conditions
go test -race ./...

# Problem: Timeout issues
go test -timeout 60s ./...

# Problem: Cache issues
go clean -testcache
go test ./...
```

### Platform-Specific Issues

#### Windows

```powershell
# PowerShell execution policy
Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser

# Path separator issues
$env:GOOS="windows"; $env:GOARCH="amd64"; go build -o hello-gopher.exe ./cmd/hello-gopher
```

#### macOS

```bash
# Code signing (for distribution)
codesign -s "Developer ID Application" hello-gopher

# Notarization (for Gatekeeper)
xcrun notarytool submit hello-gopher.zip --keychain-profile "notary-profile"
```

#### Linux

```bash
# Static linking for compatibility
CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' ./cmd/hello-gopher
```

### Performance Issues

#### Build Time Optimization

```bash
# Use build cache
export GOCACHE=/path/to/cache

# Parallel builds
go build -p 4 ./cmd/hello-gopher

# Module proxy
export GOPROXY=https://proxy.golang.org,direct
```

#### Runtime Optimization

```bash
# Profile CPU usage
go build -o hello-gopher ./cmd/hello-gopher
./hello-gopher greet --name "Profile" &
go tool pprof hello-gopher http://localhost:6060/debug/pprof/profile

# Profile memory usage
go tool pprof hello-gopher http://localhost:6060/debug/pprof/heap
```

## ⚡ Performance Optimization

### Build Optimization

```bash
# Optimize for size
go build -ldflags="-s -w" ./cmd/hello-gopher

# Optimize for speed
go build -gcflags="-B" ./cmd/hello-gopher

# Link-time optimization
go build -ldflags="-s -w -X 'main.version=optimized'" ./cmd/hello-gopher
```

### Runtime Optimization

```bash
# Set GOMAXPROCS for optimal CPU usage
export GOMAXPROCS=4

# Tune garbage collector
export GOGC=100

# Memory limit (Go 1.19+)
export GOMEMLIMIT=100MiB
```

### Profiling

```bash
# CPU profiling
go test -cpuprofile cpu.prof -bench . ./pkg/greeting
go tool pprof cpu.prof

# Memory profiling
go test -memprofile mem.prof -bench . ./pkg/greeting
go tool pprof mem.prof

# Trace analysis
go test -trace trace.out -bench . ./pkg/greeting
go tool trace trace.out
```

## 📊 Metrics and Monitoring

### Build Metrics

```bash
# Binary size analysis
ls -lh hello-gopher*

# Dependency analysis
go mod graph | wc -l

# Build time measurement
time go build ./cmd/hello-gopher
```

### Test Metrics

```bash
# Test execution time
go test -v ./... | grep -E "(PASS|FAIL|RUN)"

# Coverage metrics
go test -cover ./... | grep coverage

# Benchmark results
go test -bench=. ./... | grep Benchmark
```

### Code Quality Metrics

```bash
# Lines of code
find . -name "*.go" -not -path "./vendor/*" | xargs wc -l

# Cyclomatic complexity
gocyclo .

# Code duplication
dupl .
```

---

## 🎯 Summary

This build guide demonstrates professional Go project management with:

- **Comprehensive Testing**: Unit, integration, benchmark, and mock tests
- **Quality Assurance**: Linting, security scanning, and code formatting
- **Cross-Platform Support**: Native binaries for all major platforms
- **Automated CI/CD**: GitHub Actions with multi-stage pipeline
- **Professional Releases**: Goreleaser with Homebrew integration
- **Performance Optimization**: Profiling and optimization techniques

The Hello-Gopher CLI serves as both a functional tool and a portfolio piece showcasing Go development excellence.

For questions or issues, please refer to the [main README](README.md) or create an issue on GitHub.

**Happy Building! 🚀**