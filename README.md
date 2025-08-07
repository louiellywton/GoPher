# Hello-Gopher CLI

A friendly command-line tool that demonstrates Go development best practices.

## Features

- Greet users by name with customizable messages
- Display random Go proverbs for learning and inspiration
- Cross-platform support (Linux, macOS, Windows)
- Professional CLI experience with comprehensive help

## Installation

### Homebrew (Recommended)
```bash
brew install louiellywton/tap/hello-gopher
```

### Direct Download
Download pre-compiled binaries from the [releases page](https://github.com/louiellywton/go-portfolio/releases).

### Go Install
```bash
go install github.com/louiellywton/go-portfolio/01-hello-gopher/cmd/hello-gopher@latest
```

## Usage

### Greeting Command
```bash
# Default greeting
hello-gopher greet

# Custom name
hello-gopher greet --name Alice
hello-gopher greet -n Bob
```

### Proverb Command
```bash
# Display a random Go proverb
hello-gopher proverb
```

### Version Information
```bash
hello-gopher --version
```

## Development

This project follows Go best practices and serves as a portfolio piece demonstrating:

- Standard Go project layout
- Comprehensive testing with 80%+ coverage
- CI/CD with GitHub Actions
- Cross-platform releases with Goreleaser
- Homebrew tap integration
- Clean architecture with interfaces for testability

### Building from Source

```bash
git clone https://github.com/louiellywton/go-portfolio.git
cd go-portfolio/01-hello-gopher
go build -o hello-gopher ./cmd/hello-gopher
```

### Running Tests

```bash
go test ./... -race -cover
```

## License

MIT License - see LICENSE file for details.

## Contributing

Contributions are welcome! Please read the contributing guidelines and submit pull requests for any improvements.