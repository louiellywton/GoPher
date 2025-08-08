# PowerShell script to validate CI/CD pipeline locally on Windows
$ErrorActionPreference = "Stop"

Write-Host "🧪 Running CI/CD pipeline tests locally..." -ForegroundColor Green

Write-Host "📦 Downloading dependencies..." -ForegroundColor Yellow
go mod download
go mod verify

Write-Host "🔍 Running tests with race detection..." -ForegroundColor Yellow
go test -race -v ./...

Write-Host "📊 Running tests with coverage..." -ForegroundColor Yellow
go test -race -coverprofile coverage.out -covermode atomic ./...

Write-Host "📈 Generating coverage report..." -ForegroundColor Yellow
go tool cover -html coverage.out -o coverage.html

Write-Host "🎯 Checking coverage threshold..." -ForegroundColor Yellow
$coverageOutput = go tool cover -func coverage.out | Select-String "total"
$coverage = ($coverageOutput -split '\s+')[2] -replace '%', ''
Write-Host "Coverage: $coverage%" -ForegroundColor Cyan

if ([double]$coverage -lt 60) {
    Write-Host "❌ Coverage $coverage% is below the required 60% threshold" -ForegroundColor Red
    exit 1
}
Write-Host "✅ Coverage $coverage% meets the required 60% threshold" -ForegroundColor Green

Write-Host "🏗️ Testing build for multiple platforms..." -ForegroundColor Yellow

Write-Host "Building for Linux amd64..." -ForegroundColor Cyan
$env:GOOS = "linux"
$env:GOARCH = "amd64"
go build -o hello-gopher-linux-amd64 ./cmd/hello-gopher

Write-Host "Building for macOS amd64..." -ForegroundColor Cyan
$env:GOOS = "darwin"
$env:GOARCH = "amd64"
go build -o hello-gopher-darwin-amd64 ./cmd/hello-gopher

Write-Host "Building for Windows amd64..." -ForegroundColor Cyan
$env:GOOS = "windows"
$env:GOARCH = "amd64"
go build -o hello-gopher-windows-amd64.exe ./cmd/hello-gopher

# Reset environment variables
Remove-Item Env:GOOS -ErrorAction SilentlyContinue
Remove-Item Env:GOARCH -ErrorAction SilentlyContinue

Write-Host "🧹 Cleaning up build artifacts..." -ForegroundColor Yellow
Remove-Item hello-gopher-* -ErrorAction SilentlyContinue

Write-Host "✅ All CI/CD pipeline tests passed successfully!" -ForegroundColor Green