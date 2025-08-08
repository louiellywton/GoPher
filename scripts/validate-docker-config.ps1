# PowerShell script to validate Docker configuration
param([switch]$Verbose)

Write-Host "🔍 Validating Docker configuration for hello-gopher..." -ForegroundColor Cyan

# Check if Dockerfile exists
if (Test-Path "Dockerfile") {
    Write-Host "✓ Dockerfile exists" -ForegroundColor Green
} else {
    Write-Host "✗ Dockerfile not found" -ForegroundColor Red
    exit 1
}

# Check Dockerfile content
$dockerfileContent = Get-Content "Dockerfile" -Raw

if ($dockerfileContent -match "FROM golang:.*alpine AS builder") {
    Write-Host "✓ Build stage FROM statement found" -ForegroundColor Green
} else {
    Write-Host "✗ Build stage FROM statement missing" -ForegroundColor Red
    exit 1
}

if ($dockerfileContent -match "FROM scratch") {
    Write-Host "✓ Runtime stage FROM statement found" -ForegroundColor Green
} else {
    Write-Host "✗ Runtime stage FROM statement missing" -ForegroundColor Red
    exit 1
}

# Check .goreleaser.yaml
if (Test-Path ".goreleaser.yaml") {
    Write-Host "✓ .goreleaser.yaml exists" -ForegroundColor Green
    
    $goreleaserContent = Get-Content ".goreleaser.yaml" -Raw
    
    if ($goreleaserContent -match "dockers:") {
        Write-Host "✓ Docker configuration found in .goreleaser.yaml" -ForegroundColor Green
    } else {
        Write-Host "✗ Docker configuration missing from .goreleaser.yaml" -ForegroundColor Red
        exit 1
    }
} else {
    Write-Host "✗ .goreleaser.yaml not found" -ForegroundColor Red
    exit 1
}

# Test Go build
Write-Host "🔨 Testing Go build..." -ForegroundColor Cyan
try {
    $env:CGO_ENABLED = "0"
    $env:GOOS = "linux"
    $tempFile = [System.IO.Path]::GetTempFileName()
    
    go build -ldflags="-s -w" -o $tempFile ./cmd/hello-gopher 2>$null
    
    Write-Host "✓ Linux binary builds successfully" -ForegroundColor Green
    Remove-Item $tempFile -ErrorAction SilentlyContinue
} catch {
    Write-Host "✗ Failed to build Linux binary" -ForegroundColor Red
    exit 1
} finally {
    $env:CGO_ENABLED = $null
    $env:GOOS = $null
}

Write-Host ""
Write-Host "🎉 Docker configuration validation completed successfully!" -ForegroundColor Green