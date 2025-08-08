#!/bin/bash

# Test script to validate CI/CD pipeline locally
set -e

echo "🧪 Running CI/CD pipeline tests locally..."

echo "📦 Downloading dependencies..."
go mod download
go mod verify

echo "🔍 Running tests with race detection..."
go test -race -v ./...

echo "📊 Running tests with coverage..."
go test -race -coverprofile=coverage.out -covermode=atomic ./...

echo "📈 Generating coverage report..."
go tool cover -html=coverage.out -o coverage.html

echo "🎯 Checking coverage threshold..."
COVERAGE=$(go tool cover -func coverage.out | grep total | awk '{print $3}' | sed 's/%//')
echo "Coverage: ${COVERAGE}%"

if (( $(echo "$COVERAGE < 60" | bc -l) )); then
    echo "❌ Coverage ${COVERAGE}% is below the required 60% threshold"
    exit 1
fi
echo "✅ Coverage ${COVERAGE}% meets the required 60% threshold"

echo "🏗️ Testing build for multiple platforms..."
echo "Building for Linux amd64..."
GOOS=linux GOARCH=amd64 go build -o hello-gopher-linux-amd64 ./cmd/hello-gopher

echo "Building for macOS amd64..."
GOOS=darwin GOARCH=amd64 go build -o hello-gopher-darwin-amd64 ./cmd/hello-gopher

echo "Building for Windows amd64..."
GOOS=windows GOARCH=amd64 go build -o hello-gopher-windows-amd64.exe ./cmd/hello-gopher

echo "🧹 Cleaning up build artifacts..."
rm -f hello-gopher-*

echo "✅ All CI/CD pipeline tests passed successfully!"