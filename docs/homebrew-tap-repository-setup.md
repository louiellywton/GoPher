# Homebrew Tap Repository Setup

This document provides step-by-step instructions for creating and configuring the `homebrew-tap` repository that will host the Homebrew formula for hello-gopher.

## Prerequisites

- GitHub account (louiellywton)
- Git installed locally
- Access to create repositories in the GitHub account

## Step 1: Create the Repository

1. Go to GitHub and create a new repository:
   - Repository name: `homebrew-tap`
   - Description: "Homebrew tap for louiellywton's tools"
   - Public repository (required for Homebrew taps)
   - Initialize with README: No (we'll create our own)

2. Clone the repository locally:
   ```bash
   git clone https://github.com/louiellywton/homebrew-tap.git
   cd homebrew-tap
   ```

## Step 2: Create Repository Structure

Create the necessary directory structure and files:

```bash
# Create the Formula directory (required by Homebrew)
mkdir Formula

# Create README.md
cat > README.md << 'EOF'
# Homebrew Tap for louiellywton

This tap contains Homebrew formulas for tools developed by louiellywton.

## Installation

First, add the tap:
```bash
brew tap louiellywton/tap
```

Then install the desired package:
```bash
brew install hello-gopher
```

## Available Formulas

### hello-gopher

A friendly CLI tool for Go enthusiasts that demonstrates best practices.

**Features:**
- Greet users by name with customizable messages
- Display random Go proverbs for learning and inspiration
- Cross-platform support (Linux, macOS, Windows)
- Professional CLI experience with comprehensive help

**Installation:**
```bash
brew install louiellywton/tap/hello-gopher
```

**Usage:**
```bash
# Default greeting
hello-gopher greet

# Custom name
hello-gopher greet --name Alice

# Display a random Go proverb
hello-gopher proverb

# Show version information
hello-gopher --version
```

## Maintenance

This tap is automatically maintained by [Goreleaser](https://goreleaser.com/). Formulas are updated automatically when new releases are published to the source repositories.

## Issues and Support

If you encounter any issues with the formulas in this tap, please report them in the respective source repositories:

- [hello-gopher issues](https://github.com/louiellywton/go-portfolio/issues)

## Contributing

Contributions are welcome! However, please note that formulas are automatically generated and updated by Goreleaser. Manual changes may be overwritten during the next release.

For feature requests or bug reports, please use the source repository issue trackers.
EOF

# Create .gitignore
cat > .gitignore << 'EOF'
# macOS
.DS_Store
.DS_Store?
._*
.Spotlight-V100
.Trashes
ehthumbs.db
Thumbs.db

# Editor files
*.swp
*.swo
*~
.vscode/
.idea/

# Temporary files
*.tmp
*.temp
EOF

# Create initial commit
git add .
git commit -m "Initial tap repository setup

- Add README with installation instructions
- Create Formula directory for Homebrew formulas
- Add .gitignore for common temporary files"

git push origin main
```

## Step 3: Configure Repository Settings

1. **Branch Protection (Optional but Recommended):**
   - Go to Settings → Branches
   - Add rule for `main` branch
   - Enable "Restrict pushes that create files"
   - Enable "Require status checks to pass before merging"

2. **Repository Settings:**
   - Go to Settings → General
   - Ensure "Issues" is enabled for user feedback
   - Set "Merge button" preferences as desired
   - Consider enabling "Automatically delete head branches"

3. **Access Control:**
   - Go to Settings → Manage access
   - Ensure appropriate collaborators have access
   - Consider adding a GitHub token for Goreleaser with appropriate permissions

## Step 4: Verify Goreleaser Configuration

Ensure the Goreleaser configuration in the main project (`.goreleaser.yaml`) correctly references the tap repository:

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

## Step 5: Test the Setup

### Manual Testing

1. **Add the tap locally:**
   ```bash
   brew tap louiellywton/tap
   ```

2. **Verify the tap was added:**
   ```bash
   brew tap | grep louiellywton
   ```

3. **Check tap information:**
   ```bash
   brew tap-info louiellywton/tap
   ```

### Automated Testing (Optional)

Create a GitHub Actions workflow in the tap repository to test formulas:

```bash
mkdir -p .github/workflows

cat > .github/workflows/test.yml << 'EOF'
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
      - name: Checkout
        uses: actions/checkout@v4
      
      - name: Set up Homebrew
        id: set-up-homebrew
        uses: Homebrew/actions/setup-homebrew@master
      
      - name: Test tap
        run: |
          # Test tap syntax
          brew tap-new louiellywton/tap --no-git || true
          brew audit --tap=louiellywton/tap
          
          # Test formulas if they exist
          if ls Formula/*.rb 1> /dev/null 2>&1; then
            for formula in Formula/*.rb; do
              echo "Testing formula: $formula"
              brew audit --strict "$formula" || true
            done
          else
            echo "No formulas found to test"
          fi
EOF

git add .github/workflows/test.yml
git commit -m "Add GitHub Actions workflow for formula testing"
git push origin main
```

## Step 6: Release Integration

When you create a release in the main project:

1. **Goreleaser will automatically:**
   - Build cross-platform binaries
   - Create a GitHub release
   - Generate the Homebrew formula
   - Commit the formula to the tap repository
   - Create a pull request (if configured) or push directly

2. **The formula will be available at:**
   - File: `Formula/hello-gopher.rb`
   - Installation: `brew install louiellywton/tap/hello-gopher`

## Troubleshooting

### Common Issues

1. **Permission Denied:**
   - Ensure the GitHub token has appropriate permissions
   - Check repository access settings

2. **Formula Not Found:**
   - Verify the tap was added correctly: `brew tap`
   - Check if the formula file exists in the repository
   - Update Homebrew: `brew update`

3. **Installation Fails:**
   - Check the formula syntax: `brew audit louiellywton/tap/hello-gopher`
   - Verify download URLs are accessible
   - Check SHA256 checksums match

### Debugging Commands

```bash
# List all taps
brew tap

# Show tap information
brew tap-info louiellywton/tap

# Audit a specific formula
brew audit --strict louiellywton/tap/hello-gopher

# Test formula installation
brew install --build-from-source louiellywton/tap/hello-gopher

# Show formula contents
brew cat louiellywton/tap/hello-gopher

# Uninstall and clean up
brew uninstall hello-gopher
brew untap louiellywton/tap
```

## Security Considerations

1. **Repository Access:**
   - Keep the repository public (required for Homebrew)
   - Limit write access to trusted contributors
   - Use branch protection rules

2. **Token Security:**
   - Use a GitHub token with minimal required permissions
   - Store tokens securely in GitHub Secrets
   - Regularly rotate tokens

3. **Formula Security:**
   - Verify checksums are correctly generated
   - Monitor for unauthorized changes
   - Use signed releases when possible

## Maintenance

1. **Regular Tasks:**
   - Monitor the repository for issues
   - Update documentation as needed
   - Review and merge any manual pull requests
   - Keep GitHub Actions workflows updated

2. **Monitoring:**
   - Watch for failed releases or formula updates
   - Monitor user feedback and issues
   - Check formula audit results regularly

This setup ensures that your Homebrew tap is properly configured and ready for automatic updates via Goreleaser.