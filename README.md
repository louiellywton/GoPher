# Hello-Gopher CLI

[![CI/CD Pipeline](https://github.com/louiellywton/go-portfolio/actions/workflows/ci.yml/badge.svg)](https://github.com/louiellywton/go-portfolio/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/louiellywton/go-portfolio/01-hello-gopher)](https://goreportcard.com/report/github.com/louiellywton/go-portfolio/01-hello-gopher)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A friendly command-line tool that demonstrates Go development best practices. Hello-Gopher provides greeting functionality and displays random Go proverbs, serving as a portfolio piece that showcases idiomatic Go code, comprehensive testing, CI/CD integration, and professional distribution through Homebrew.

## ✨ Features

- **Personalized Greetings**: Greet users by name with customizable messages
- **Go Proverbs**: Display random Go proverbs for learning and inspiration
- **Cross-Platform**: Native support for Linux, macOS, and Windows (amd64 & arm64)
- **Professional CLI**: Comprehensive help system with examples and error handling
- **Zero Dependencies**: Self-contained binary with embedded resources
- **Fast & Lightweight**: Minimal resource usage with quick startup time

## 🚀 Installation

### Homebrew (Recommended)

The easiest way to install on macOS and Linux:

```bash
brew install louiellywton/tap/hello-gopher
```

### Direct Download

Download pre-compiled binaries for your platform from the [releases page](https://github.com/louiellywton/go-portfolio/releases):

- **Linux**: `hello-gopher_Linux_x86_64.tar.gz` or `hello-gopher_Linux_arm64.tar.gz`
- **macOS**: `hello-gopher_Darwin_x86_64.tar.gz` or `hello-gopher_Darwin_arm64.tar.gz`
- **Windows**: `hello-gopher_Windows_x86_64.zip` or `hello-gopher_Windows_arm64.zip`

Extract the archive and move the binary to your PATH.

### Go Install

If you have Go installed:

```bash
go install github.com/louiellywton/go-portfolio/01-hello-gopher/cmd/hello-gopher@latest
```

### Docker

Run hello-gopher in a container without installing anything locally:

```bash
# Pull the latest image
docker pull ghcr.io/louiellywton/hello-gopher:latest

# Run with default greeting
docker run --rm ghcr.io/louiellywton/hello-gopher:latest greet

# Run with custom name
docker run --rm ghcr.io/louiellywton/hello-gopher:latest greet --name Docker

# Display a random proverb
docker run --rm ghcr.io/louiellywton/hello-gopher:latest proverb

# Show version information
docker run --rm ghcr.io/louiellywton/hello-gopher:latest --version

# Show help
docker run --rm ghcr.io/louiellywton/hello-gopher:latest --help
```

**Available Tags:**
- `latest` - Latest stable release (multi-arch: amd64/arm64)
- `v1.x.x` - Specific version tags
- `v1.x` - Major.minor version tags

## 📖 Usage

### Getting Help

```bash
# Show general help
hello-gopher --help

# Show help for specific commands
hello-gopher greet --help
hello-gopher proverb --help
```

**Output:**
```
Hello-Gopher is a friendly command-line tool that demonstrates Go development best practices.
It provides greeting functionality and displays random Go proverbs, serving as a portfolio piece
that showcases idiomatic Go code, comprehensive testing, and professional distribution.

Examples:
  hello-gopher greet                    # Greet the default gopher
  hello-gopher greet --name Alice       # Greet Alice
  hello-gopher greet -n Bob             # Greet Bob (short flag)
  hello-gopher proverb                  # Display a random Go proverb
  hello-gopher --version                # Show version information

Available Commands:
  greet       Greet a gopher by name
  help        Help about any command
  proverb     Display a random Go proverb

Flags:
  -h, --help      help for hello-gopher
  -v, --version   version for hello-gopher
```

### Greeting Command

```bash
# Default greeting
hello-gopher greet
```
**Output:** `Hello, Gopher!`

```bash
# Custom name with long flag
hello-gopher greet --name Alice
```
**Output:** `Hello, Alice!`

```bash
# Custom name with short flag
hello-gopher greet -n Bob
```
**Output:** `Hello, Bob!`

### Proverb Command

```bash
# Display a random Go proverb
hello-gopher proverb
```
**Sample Output:** `Don't communicate by sharing memory, share memory by communicating.`

Each execution shows a different proverb from a curated collection of 50+ Go programming wisdom and best practices.

### Version Information

```bash
hello-gopher --version
```
**Sample Output:**
```
hello-gopher version v1.0.0
Build date: 2024-01-15T10:30:00Z
Git commit: abc123def456
Go version: go1.22.0
OS/Arch: linux/amd64
```

## 🏗️ Development

This project serves as a comprehensive example of Go development best practices, demonstrating:

### Architecture & Design
- **Standard Go Project Layout**: Organized following community conventions
- **Clean Architecture**: Interfaces for testability and maintainability
- **Dependency Injection**: Mockable components for isolated testing
- **Error Handling**: Custom error types with helpful user guidance

### Testing & Quality
- **Comprehensive Test Coverage**: 80%+ coverage with race condition detection
- **Table-Driven Tests**: Idiomatic Go testing patterns
- **Benchmark Tests**: Performance validation and optimization
- **Mock Testing**: Interface-based testing for clean architecture
- **Integration Tests**: End-to-end command testing

### CI/CD & Distribution
- **GitHub Actions**: Multi-platform testing and automated releases
- **Cross-Platform Builds**: Native binaries for all major platforms
- **Goreleaser**: Professional release automation with checksums
- **Homebrew Integration**: Automated tap updates for easy installation
- **Security Scanning**: Vulnerability detection and dependency analysis

### Prerequisites

- Go 1.21 or later
- Git

### Building from Source

```bash
# Clone the repository
git clone https://github.com/louiellywton/go-portfolio.git
cd go-portfolio/01-hello-gopher

# Download dependencies
go mod download

# Build the binary
go build -o hello-gopher ./cmd/hello-gopher

# Run the binary
./hello-gopher greet --name Developer
```

### Development Commands

```bash
# Run tests with coverage
go test ./... -race -cover

# Run tests with verbose output
go test -v ./...

# Run benchmarks
go test -bench=. -benchmem ./...

# Generate coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html

# Run linting (requires golangci-lint)
golangci-lint run

# Run security scan (requires gosec)
gosec ./...

# Check for vulnerabilities
govulncheck ./...
```

### Docker Development

```bash
# Build Docker image locally
docker build -t hello-gopher:dev .

# Run the local image
docker run --rm hello-gopher:dev greet --name Developer

# Test the container
docker run --rm hello-gopher:dev --version

# Build multi-arch image (requires buildx)
docker buildx build --platform linux/amd64,linux/arm64 -t hello-gopher:multi-arch .
```

### Project Structure

```
01-hello-gopher/
├── cmd/
│   └── hello-gopher/           # Application entry point
│       ├── main.go             # Main function
│       └── cmd/                # Cobra commands
│           ├── root.go         # Root command setup
│           ├── greet.go        # Greet command
│           ├── proverb.go      # Proverb command
│           ├── version.go      # Version command
│           └── errors.go       # Error handling
├── pkg/
│   └── greeting/               # Core business logic
│       ├── greeting.go         # Greeting functionality
│       ├── proverb.go         # Proverb functionality
│       └── *_test.go          # Test files
├── scripts/
│   └── proverb.txt            # Embedded proverb data
├── .github/
│   └── workflows/
│       └── ci.yml             # CI/CD pipeline
├── .goreleaser.yaml           # Release configuration
├── BUILD.md                   # Detailed build guide
└── README.md                  # This file
```

## 🤝 Contributing

Contributions are welcome! This project follows standard Go contribution practices.

### Getting Started

1. **Fork the repository** on GitHub
2. **Clone your fork** locally:
   ```bash
   git clone https://github.com/your-username/go-portfolio.git
   cd go-portfolio/01-hello-gopher
   ```
3. **Create a feature branch**:
   ```bash
   git checkout -b feature/your-feature-name
   ```
4. **Make your changes** and add tests
5. **Run the test suite**:
   ```bash
   go test ./... -race -cover
   ```
6. **Run linting**:
   ```bash
   golangci-lint run
   ```
7. **Commit your changes**:
   ```bash
   git commit -am "Add your feature description"
   ```
8. **Push to your fork**:
   ```bash
   git push origin feature/your-feature-name
   ```
9. **Create a Pull Request** on GitHub

### Development Guidelines

- **Code Style**: Follow `gofmt` and `golangci-lint` standards
- **Testing**: Maintain 80%+ test coverage for new code
- **Documentation**: Add Go doc comments for public APIs
- **Commits**: Use clear, descriptive commit messages
- **Issues**: Reference issue numbers in commits when applicable

### Code of Conduct

This project follows the [Go Community Code of Conduct](https://golang.org/conduct). Please be respectful and inclusive in all interactions.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- **Go Team**: For creating an amazing programming language
- **Cobra**: For the excellent CLI framework
- **Goreleaser**: For simplifying cross-platform releases
- **Go Community**: For the wealth of best practices and conventions

## 📊 Project Stats

- **Language**: Go 1.22+
- **Dependencies**: Minimal (only Cobra for CLI)
- **Test Coverage**: 80%+
- **Platforms**: Linux, macOS, Windows (amd64 & arm64)
- **Binary Size**: ~8MB (compressed: ~3MB)
- **Startup Time**: <10ms

---

**Built with ❤️ and Go** | [Portfolio](https://github.com/louiellywton/go-portfolio) | [Issues](https://github.com/louiellywton/go-portfolio/issues) | [Releases](https://github.com/louiellywton/go-portfolio/releases)
