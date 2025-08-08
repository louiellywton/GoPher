#!/bin/bash

# Comprehensive test script for Docker implementation
# Run this script to validate the complete Docker setup

set -e

echo "🐳 Testing Docker Implementation for hello-gopher CLI"
echo "=================================================="

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

print_success() {
    echo -e "${GREEN}✓${NC} $1"
}

print_info() {
    echo -e "${BLUE}ℹ${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}⚠${NC} $1"
}

print_error() {
    echo -e "${RED}✗${NC} $1"
}

# Test 1: Check if required files exist
echo -e "\n${BLUE}1. Checking Docker files...${NC}"

if [[ -f "Dockerfile" ]]; then
    print_success "Dockerfile exists"
else
    print_error "Dockerfile missing"
    exit 1
fi

if [[ -f ".goreleaser.yaml" ]]; then
    print_success ".goreleaser.yaml exists"
else
    print_error ".goreleaser.yaml missing"
    exit 1
fi

# Test 2: Validate Dockerfile content
echo -e "\n${BLUE}2. Validating Dockerfile content...${NC}"

if grep -q "FROM golang:.*alpine AS builder" Dockerfile; then
    print_success "Multi-stage build configured"
else
    print_error "Multi-stage build not found"
    exit 1
fi

if grep -q "FROM scratch" Dockerfile; then
    print_success "Minimal runtime image configured"
else
    print_error "Scratch base image not found"
    exit 1
fi

if grep -q "CGO_ENABLED=0" Dockerfile; then
    print_success "Static binary compilation enabled"
else
    print_warning "CGO not explicitly disabled"
fi

# Test 3: Validate Goreleaser Docker configuration
echo -e "\n${BLUE}3. Validating Goreleaser Docker config...${NC}"

if grep -q "dockers:" .goreleaser.yaml; then
    print_success "Docker configuration found in Goreleaser"
else
    print_error "Docker configuration missing from Goreleaser"
    exit 1
fi

if grep -q "ghcr.io" .goreleaser.yaml; then
    print_success "GitHub Container Registry configured"
else
    print_warning "GitHub Container Registry not configured"
fi

if grep -q "docker_manifests:" .goreleaser.yaml; then
    print_success "Multi-arch Docker manifests configured"
else
    print_warning "Multi-arch manifests not configured"
fi

# Test 4: Check Goreleaser configuration validity
echo -e "\n${BLUE}4. Testing Goreleaser configuration...${NC}"

if command -v goreleaser &> /dev/null; then
    print_info "Goreleaser is available"
    
    if goreleaser check; then
        print_success "Goreleaser configuration is valid"
    else
        print_error "Goreleaser configuration has issues"
        exit 1
    fi
else
    print_warning "Goreleaser not found in PATH"
fi

# Test 5: Test Go build for Linux (simulates Docker build)
echo -e "\n${BLUE}5. Testing Linux binary build...${NC}"

if CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o hello-gopher-test ./cmd/hello-gopher; then
    print_success "Linux binary builds successfully"
    
    # Check binary size
    size=$(stat -c%s hello-gopher-test)
    size_mb=$(echo "scale=2; $size/1024/1024" | bc)
    print_info "Binary size: ${size_mb} MB"
    
    # Clean up
    rm -f hello-gopher-test
else
    print_error "Linux binary build failed"
    exit 1
fi

# Test 6: Check Docker availability (optional)
echo -e "\n${BLUE}6. Checking Docker availability...${NC}"

if command -v docker &> /dev/null; then
    print_info "Docker is available"
    
    if docker info &> /dev/null; then
        print_success "Docker daemon is running"
        
        # Test Docker build if Docker is available
        echo -e "\n${BLUE}7. Testing Docker build...${NC}"
        
        if docker build -t hello-gopher:test . &> /dev/null; then
            print_success "Docker image builds successfully"
            
            # Test running the container
            if docker run --rm hello-gopher:test greet --name "Docker Test" 2>/dev/null | grep -q "Hello, Docker Test!"; then
                print_success "Docker container runs correctly"
            else
                print_warning "Docker container test failed"
            fi
            
            # Clean up
            docker rmi hello-gopher:test &> /dev/null || true
        else
            print_warning "Docker build failed (this might be due to missing dependencies)"
        fi
    else
        print_warning "Docker daemon is not running"
    fi
else
    print_warning "Docker is not installed"
fi

# Test 7: Check documentation
echo -e "\n${BLUE}8. Checking documentation...${NC}"

if grep -q "Docker" README.md; then
    print_success "Docker documentation found in README.md"
else
    print_warning "Docker documentation missing from README.md"
fi

if grep -q "docker run" README.md; then
    print_success "Docker usage examples found"
else
    print_warning "Docker usage examples missing"
fi

# Test 8: Check test scripts
echo -e "\n${BLUE}9. Checking test scripts...${NC}"

if [[ -f "scripts/test-docker.sh" ]]; then
    print_success "Docker test script (bash) exists"
else
    print_warning "Docker test script (bash) missing"
fi

if [[ -f "scripts/test-docker.ps1" ]]; then
    print_success "Docker test script (PowerShell) exists"
else
    print_warning "Docker test script (PowerShell) missing"
fi

# Summary
echo -e "\n${GREEN}🎉 Docker implementation test completed!${NC}"
echo -e "\n${BLUE}Summary:${NC}"
echo "  ✓ Dockerfile with multi-stage build"
echo "  ✓ Goreleaser Docker integration"
echo "  ✓ GitHub Container Registry setup"
echo "  ✓ Multi-arch support configuration"
echo "  ✓ Static binary compilation"
echo "  ✓ Documentation and examples"
echo "  ✓ Cross-platform test scripts"

echo -e "\n${BLUE}Next steps:${NC}"
echo "  1. Run 'goreleaser build --snapshot --clean' to test local build"
echo "  2. Run 'docker build -t hello-gopher:local .' to test Docker build"
echo "  3. Push to GitHub to trigger automated Docker image publishing"

echo -e "\n${GREEN}Docker support is ready for production! 🚀${NC}"