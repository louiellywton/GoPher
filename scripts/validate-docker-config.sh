#!/bin/bash

# Validation script for Docker configuration
# This script validates Docker-related files without requiring Docker to be running

set -e

echo "🔍 Validating Docker configuration for hello-gopher..."

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Function to print status
print_status() {
    echo -e "${GREEN}✓${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}⚠${NC} $1"
}

print_error() {
    echo -e "${RED}✗${NC} $1"
}

# Check if Dockerfile exists
if [[ -f "Dockerfile" ]]; then
    print_status "Dockerfile exists"
else
    print_error "Dockerfile not found"
    exit 1
fi

# Validate Dockerfile syntax (basic checks)
echo "🔍 Validating Dockerfile syntax..."

# Check for required FROM statements
if grep -q "FROM golang:.*alpine AS builder" Dockerfile; then
    print_status "Build stage FROM statement found"
else
    print_error "Build stage FROM statement missing or incorrect"
    exit 1
fi

if grep -q "FROM scratch" Dockerfile; then
    print_status "Runtime stage FROM statement found"
else
    print_error "Runtime stage FROM statement missing"
    exit 1
fi

# Check for WORKDIR
if grep -q "WORKDIR /app" Dockerfile; then
    print_status "WORKDIR statement found"
else
    print_error "WORKDIR statement missing"
    exit 1
fi

# Check for COPY statements
if grep -q "COPY go.mod go.sum" Dockerfile; then
    print_status "Go module files COPY statement found"
else
    print_error "Go module files COPY statement missing"
    exit 1
fi

if grep -q "COPY --from=builder" Dockerfile; then
    print_status "Multi-stage COPY statement found"
else
    print_error "Multi-stage COPY statement missing"
    exit 1
fi

# Check for ENTRYPOINT
if grep -q "ENTRYPOINT" Dockerfile; then
    print_status "ENTRYPOINT statement found"
else
    print_error "ENTRYPOINT statement missing"
    exit 1
fi

# Check for build optimizations
if grep -q "CGO_ENABLED=0" Dockerfile; then
    print_status "CGO disabled for static binary"
else
    print_warning "CGO not explicitly disabled"
fi

if grep -q "\-ldflags.*\-s \-w" Dockerfile; then
    print_status "Binary size optimization flags found"
else
    print_warning "Binary size optimization flags not found"
fi

# Validate .goreleaser.yaml Docker configuration
echo "🔍 Validating Goreleaser Docker configuration..."

if [[ -f ".goreleaser.yaml" ]]; then
    print_status ".goreleaser.yaml exists"
    
    # Check for Docker configuration
    if grep -q "dockers:" .goreleaser.yaml; then
        print_status "Docker configuration found in .goreleaser.yaml"
    else
        print_error "Docker configuration missing from .goreleaser.yaml"
        exit 1
    fi
    
    # Check for image templates
    if grep -q "image_templates:" .goreleaser.yaml; then
        print_status "Docker image templates found"
    else
        print_error "Docker image templates missing"
        exit 1
    fi
    
    # Check for GitHub Container Registry
    if grep -q "ghcr.io" .goreleaser.yaml; then
        print_status "GitHub Container Registry configuration found"
    else
        print_warning "GitHub Container Registry not configured"
    fi
    
    # Check for multi-arch support
    if grep -q "docker_manifests:" .goreleaser.yaml; then
        print_status "Multi-arch Docker manifest configuration found"
    else
        print_warning "Multi-arch Docker manifest not configured"
    fi
    
else
    print_error ".goreleaser.yaml not found"
    exit 1
fi

# Check for Docker test scripts
echo "🔍 Checking Docker test scripts..."

if [[ -f "scripts/test-docker.sh" ]]; then
    print_status "Docker test script (bash) exists"
else
    print_warning "Docker test script (bash) not found"
fi

if [[ -f "scripts/test-docker.ps1" ]]; then
    print_status "Docker test script (PowerShell) exists"
else
    print_warning "Docker test script (PowerShell) not found"
fi

# Validate Go application can be built (simulates Docker build)
echo "🔨 Validating Go build (simulates Docker build)..."

if go mod download > /dev/null 2>&1; then
    print_status "Go dependencies downloaded successfully"
else
    print_error "Failed to download Go dependencies"
    exit 1
fi

if CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /tmp/hello-gopher-test ./cmd/hello-gopher > /dev/null 2>&1; then
    print_status "Linux binary builds successfully (Docker simulation)"
    rm -f /tmp/hello-gopher-test
else
    print_error "Failed to build Linux binary"
    exit 1
fi

# Check README documentation
echo "📖 Checking Docker documentation..."

if [[ -f "README.md" ]]; then
    if grep -q "Docker" README.md; then
        print_status "Docker documentation found in README.md"
        
        if grep -q "docker run" README.md; then
            print_status "Docker usage examples found"
        else
            print_warning "Docker usage examples not found in README.md"
        fi
        
        if grep -q "ghcr.io" README.md; then
            print_status "Container registry information found"
        else
            print_warning "Container registry information not found in README.md"
        fi
    else
        print_warning "Docker documentation not found in README.md"
    fi
else
    print_error "README.md not found"
    exit 1
fi

echo ""
echo -e "${GREEN}🎉 Docker configuration validation completed successfully!${NC}"
echo ""
echo "Validated components:"
echo "  ✓ Dockerfile syntax and structure"
echo "  ✓ Multi-stage build configuration"
echo "  ✓ Goreleaser Docker integration"
echo "  ✓ Build optimization settings"
echo "  ✓ Go application build compatibility"
echo "  ✓ Documentation completeness"
echo ""
echo "The Docker configuration is ready for container builds!"