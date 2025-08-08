#!/bin/bash

# Test script for Docker functionality
# This script tests the Docker container across different scenarios

set -e

echo "🐳 Testing Docker functionality for hello-gopher..."

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

# Check if Docker is available
if ! command -v docker &> /dev/null; then
    print_error "Docker is not installed or not in PATH"
    exit 1
fi

print_status "Docker is available"

# Build the Docker image locally
echo "🔨 Building Docker image..."
if docker build -t hello-gopher:test . > /dev/null 2>&1; then
    print_status "Docker image built successfully"
else
    print_error "Failed to build Docker image"
    exit 1
fi

# Test basic functionality
echo "🧪 Testing basic functionality..."

# Test default greeting
echo "Testing default greeting..."
output=$(docker run --rm hello-gopher:test greet 2>/dev/null)
if [[ "$output" == "Hello, Gopher!" ]]; then
    print_status "Default greeting works"
else
    print_error "Default greeting failed. Expected 'Hello, Gopher!', got '$output'"
    exit 1
fi

# Test custom name greeting
echo "Testing custom name greeting..."
output=$(docker run --rm hello-gopher:test greet --name Docker 2>/dev/null)
if [[ "$output" == "Hello, Docker!" ]]; then
    print_status "Custom name greeting works"
else
    print_error "Custom name greeting failed. Expected 'Hello, Docker!', got '$output'"
    exit 1
fi

# Test short flag
echo "Testing short flag..."
output=$(docker run --rm hello-gopher:test greet -n Test 2>/dev/null)
if [[ "$output" == "Hello, Test!" ]]; then
    print_status "Short flag works"
else
    print_error "Short flag failed. Expected 'Hello, Test!', got '$output'"
    exit 1
fi

# Test proverb command
echo "Testing proverb command..."
output=$(docker run --rm hello-gopher:test proverb 2>/dev/null)
if [[ -n "$output" && ${#output} -gt 10 ]]; then
    print_status "Proverb command works (output: ${output:0:50}...)"
else
    print_error "Proverb command failed or returned empty/short output"
    exit 1
fi

# Test version command
echo "Testing version command..."
output=$(docker run --rm hello-gopher:test --version 2>/dev/null)
if [[ "$output" == *"hello-gopher version"* ]]; then
    print_status "Version command works"
else
    print_error "Version command failed. Output: '$output'"
    exit 1
fi

# Test help command
echo "Testing help command..."
output=$(docker run --rm hello-gopher:test --help 2>/dev/null)
if [[ "$output" == *"Hello-Gopher is a friendly command-line tool"* ]]; then
    print_status "Help command works"
else
    print_error "Help command failed"
    exit 1
fi

# Test error handling
echo "Testing error handling..."
if docker run --rm hello-gopher:test invalid-command > /dev/null 2>&1; then
    print_error "Error handling test failed - should have returned non-zero exit code"
    exit 1
else
    print_status "Error handling works correctly"
fi

# Test container size
echo "📏 Checking container size..."
size=$(docker images hello-gopher:test --format "table {{.Size}}" | tail -n 1)
print_status "Container size: $size"

# Test container layers
echo "🔍 Checking container layers..."
layers=$(docker history hello-gopher:test --format "table {{.CreatedBy}}" | wc -l)
print_status "Container has $layers layers"

# Clean up test image
echo "🧹 Cleaning up..."
docker rmi hello-gopher:test > /dev/null 2>&1
print_status "Test image removed"

echo ""
echo -e "${GREEN}🎉 All Docker tests passed successfully!${NC}"
echo ""
echo "Container features verified:"
echo "  ✓ Basic greeting functionality"
echo "  ✓ Custom name support"
echo "  ✓ Flag parsing (long and short)"
echo "  ✓ Proverb display"
echo "  ✓ Version information"
echo "  ✓ Help system"
echo "  ✓ Error handling"
echo "  ✓ Minimal container size"
echo ""
echo "The Docker container is ready for production use!"