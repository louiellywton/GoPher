# CI/CD Pipeline Documentation

## Overview

The Hello-Gopher CLI project uses GitHub Actions for continuous integration and continuous deployment (CI/CD). The pipeline ensures code quality, security, and automated releases across multiple platforms.

## Pipeline Components

### 1. Main CI/CD Workflow (`.github/workflows/ci.yml`)

The main workflow includes the following jobs:

#### Test Job
- **Platforms**: Ubuntu, macOS, Windows
- **Go Versions**: 1.21, 1.22
- **Features**:
  - Race condition detection
  - Code coverage reporting
  - Dependency verification
  - Codecov integration

#### Lint Job
- **Tool**: golangci-lint
- **Configuration**: `.golangci.yml`
- **Features**:
  - Code style enforcement
  - Static analysis
  - Best practices validation

#### Security Job
- **Tools**: 
  - Gosec security scanner
  - govulncheck vulnerability scanner
- **Features**:
  - Security vulnerability detection
  - Dependency security scanning

#### Dependency Scan Job
- **Tool**: Trivy vulnerability scanner
- **Features**:
  - File system scanning
  - SARIF report generation
  - GitHub Security tab integration

#### Build Job
- **Platforms**: Linux, macOS, Windows
- **Architectures**: amd64, arm64
- **Features**:
  - Cross-platform binary generation
  - Version information injection
  - Build artifact upload

#### Coverage Report Job
- **Features**:
  - HTML coverage report generation
  - Coverage threshold validation (60%)
  - Coverage artifact upload

#### Benchmark Job
- **Features**:
  - Performance benchmark execution
  - Memory usage analysis

#### Release Job (on release events)
- **Tool**: GoReleaser
- **Features**:
  - Automated release creation
  - Cross-platform binary distribution
  - Homebrew tap integration

### 2. CodeQL Security Analysis (`.github/workflows/codeql.yml`)

- **Schedule**: Weekly on Sundays
- **Features**:
  - Advanced security analysis
  - Vulnerability detection
  - Security advisory integration

### 3. Dependabot Configuration (`.github/dependabot.yml`)

- **Go Modules**: Weekly updates on Mondays
- **GitHub Actions**: Weekly updates on Mondays
- **Features**:
  - Automated dependency updates
  - Security vulnerability patching
  - Pull request automation

## Configuration Files

### golangci-lint Configuration (`.golangci.yml`)

Comprehensive linting configuration with:
- 30+ enabled linters
- Custom rules for test files
- Performance and style checks
- Security vulnerability detection

### Local Testing Scripts

#### PowerShell Script (`scripts/test-ci.ps1`)
```powershell
powershell -ExecutionPolicy Bypass -File scripts/test-ci.ps1
```

#### Bash Script (`scripts/test-ci.sh`)
```bash
chmod +x scripts/test-ci.sh
./scripts/test-ci.sh
```

## Pipeline Triggers

### Push Events
- Branches: `main`, `develop`
- All jobs execute

### Pull Request Events
- Target branch: `main`
- All jobs execute except release

### Release Events
- Event type: `published`
- Triggers release job with GoReleaser

### Scheduled Events
- CodeQL analysis: Weekly on Sundays at 1:30 AM UTC

## Environment Variables and Secrets

### Required Secrets
- `GITHUB_TOKEN`: Automatically provided by GitHub
- `HOMEBREW_TAP_GITHUB_TOKEN`: Personal access token for Homebrew tap updates

### Environment Variables
- `GO_VERSION_MATRIX`: JSON array of supported Go versions
- Build-time variables injected via ldflags:
  - `main.version`: Git tag or branch name
  - `main.buildDate`: ISO 8601 build timestamp
  - `main.gitCommit`: Git commit SHA

## Quality Gates

### Test Coverage
- **Minimum**: 60% overall coverage
- **Target**: 80% for core business logic
- **Reporting**: Codecov integration with PR comments

### Security Scanning
- **Gosec**: Go security analyzer
- **govulncheck**: Go vulnerability database
- **Trivy**: Container and filesystem scanner
- **CodeQL**: GitHub's semantic code analysis

### Code Quality
- **golangci-lint**: 30+ linters enabled
- **Race detection**: All tests run with `-race` flag
- **Dependency verification**: `go mod verify`

## Artifacts

### Build Artifacts
- Cross-platform binaries (Linux, macOS, Windows)
- Architecture support (amd64, arm64)
- Compressed archives with checksums

### Coverage Artifacts
- `coverage.out`: Machine-readable coverage data
- `coverage.html`: Human-readable coverage report

### Security Artifacts
- SARIF reports uploaded to GitHub Security tab
- Vulnerability scan results

## Performance Monitoring

### Benchmarks
- Greeting function performance
- Proverb loading performance
- Memory allocation tracking

### Build Performance
- Parallel job execution
- Dependency caching
- Artifact caching

## Troubleshooting

### Common Issues

1. **Coverage Below Threshold**
   - Add more unit tests
   - Test edge cases
   - Mock external dependencies

2. **Security Vulnerabilities**
   - Update dependencies with `go get -u`
   - Review Dependabot PRs
   - Check GitHub Security advisories

3. **Lint Failures**
   - Run `golangci-lint run` locally
   - Fix code style issues
   - Update `.golangci.yml` if needed

4. **Build Failures**
   - Check Go version compatibility
   - Verify cross-platform code
   - Test locally with build scripts

### Local Development

Run the complete CI pipeline locally:

```bash
# Windows
powershell -ExecutionPolicy Bypass -File scripts/test-ci.ps1

# Unix/Linux/macOS
chmod +x scripts/test-ci.sh
./scripts/test-ci.sh
```

## Metrics and Monitoring

### Success Metrics
- ✅ All tests pass across platforms
- ✅ Coverage above 60% threshold
- ✅ No security vulnerabilities
- ✅ Clean lint results
- ✅ Successful cross-platform builds

### Performance Metrics
- Build time: < 5 minutes
- Test execution: < 2 minutes per platform
- Coverage generation: < 30 seconds

## Future Enhancements

### Planned Improvements
- [ ] Increase coverage threshold to 80%
- [ ] Add integration tests with Docker
- [ ] Implement canary deployments
- [ ] Add performance regression detection
- [ ] Integrate with external monitoring tools

### Advanced Features
- [ ] Multi-environment deployments
- [ ] Blue-green deployment strategy
- [ ] Automated rollback mechanisms
- [ ] Performance benchmarking trends
- [ ] Security compliance reporting