# Homebrew Tap Setup Guide

This document describes how to set up and maintain the Homebrew tap for hello-gopher.

## Repository Structure

The Homebrew tap repository (`homebrew-tap`) should have the following structure:

```
homebrew-tap/
├── README.md
├── Formula/
│   └── hello-gopher.rb  # Generated automatically by Goreleaser
└── .github/
    └── workflows/
        └── tests.yml    # Optional: Test formula installations
```

## Initial Setup

### 1. Create the Homebrew Tap Repository

Create a new repository named `homebrew-tap` in your GitHub account:

```bash
# Create the repository on GitHub first, then clone it
git clone https://github.com/louiellywton/homebrew-tap.git
cd homebrew-tap
```

### 2. Initialize Repository Structure

```bash
# Create the Formula directory
mkdir -p Formula

# Create README
cat > README.md << 'EOF'
# Homebrew Tap for louiellywton

This tap contains Homebrew formulas for tools developed by louiellywton.

## Installation

```bash
brew tap louiellywton/tap
brew install hello-gopher
```

## Available Formulas

- **hello-gopher**: A friendly CLI tool for Go enthusiasts

## Maintenance

This tap is automatically maintained by Goreleaser. Formulas are updated automatically when new releases are published.
EOF

# Create initial commit
git add .
git commit -m "Initial tap setup"
git push origin main
```

### 3. Configure Repository Settings

1. Go to repository Settings → General
2. Set default branch to `main`
3. Enable "Allow auto-merge"
4. Configure branch protection rules if desired

## Goreleaser Configuration

The Goreleaser configuration in `.goreleaser.yaml` is already set up for Homebrew integration:

```yaml
brews:
  - repository:
      owner: louiellywton
      name: homebrew-tap
      branch: main
    
    name: hello-gopher
    homepage: https://github.com/louiellywton/go-portfolio
    description: "A friendly CLI tool for Go enthusiasts"
    license: MIT
    
    install: |
      bin.install "hello-gopher"
    
    test: |
      system "#{bin}/hello-gopher", "--version"
      system "#{bin}/hello-gopher", "greet", "--name", "Homebrew"
```

## Testing the Tap

### Local Testing

1. **Test the formula locally:**
   ```bash
   # Add the tap
   brew tap louiellywton/tap

   # Install the package
   brew install hello-gopher

   # Test the installation
   hello-gopher --version
   hello-gopher greet --name "Test User"
   hello-gopher proverb
   ```

2. **Test formula syntax:**
   ```bash
   # Audit the formula
   brew audit --strict louiellywton/tap/hello-gopher

   # Test installation from source
   brew install --build-from-source louiellywton/tap/hello-gopher
   ```

### Automated Testing

The tap can include automated testing via GitHub Actions:

```yaml
# .github/workflows/tests.yml in homebrew-tap repository
name: Test Formulas

on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]

jobs:
  test:
    runs-on: macos-latest
    steps:
      - uses: actions/checkout@v4
      
      - name: Set up Homebrew
        uses: Homebrew/actions/setup-homebrew@master
      
      - name: Test formula
        run: |
          brew test-bot --only-cleanup-before
          brew test-bot --only-setup
          brew test-bot --only-tap-syntax
          brew test-bot --only-formulae --root-url=https://github.com/louiellywton/go-portfolio/releases/download
```

## Release Process

When a new release is created:

1. Goreleaser builds cross-platform binaries
2. Creates GitHub release with assets
3. Automatically updates the Homebrew formula in the tap repository
4. Users can install the new version with `brew upgrade hello-gopher`

## Troubleshooting

### Common Issues

1. **Formula not found:**
   - Ensure the tap is added: `brew tap louiellywton/tap`
   - Check if the formula exists: `ls $(brew --repository)/Library/Taps/louiellywton/homebrew-tap/Formula/`

2. **Installation fails:**
   - Check the formula syntax: `brew audit louiellywton/tap/hello-gopher`
   - Verify the download URL is accessible
   - Check the SHA256 checksum matches

3. **Outdated formula:**
   - Update the tap: `brew update`
   - Check for new versions: `brew outdated`

### Manual Formula Update

If needed, you can manually update the formula:

```bash
# Edit the formula
brew edit louiellywton/tap/hello-gopher

# Test the changes
brew reinstall louiellywton/tap/hello-gopher
```

## Security Considerations

- The tap repository should have branch protection enabled
- Only trusted contributors should have write access
- Goreleaser should use a GitHub token with minimal required permissions
- Regular security audits of the formula and dependencies

## Maintenance

- Monitor the tap repository for issues and pull requests
- Keep Goreleaser configuration up to date
- Regularly test installations on different macOS versions
- Update documentation as needed