#!/bin/bash

# Quick test for Goreleaser Docker configuration
echo "🔍 Testing Goreleaser Docker configuration..."

# Check if goreleaser is available
if ! command -v goreleaser &> /dev/null; then
    echo "❌ Goreleaser not found. Please install it first."
    exit 1
fi

echo "✅ Goreleaser is available"

# Check configuration
echo "🔧 Checking Goreleaser configuration..."
if goreleaser check; then
    echo "✅ Goreleaser configuration is valid"
else
    echo "❌ Goreleaser configuration has issues"
    exit 1
fi

# Test build without Docker (snapshot)
echo "🔨 Testing snapshot build..."
if goreleaser build --snapshot --clean --skip=validate; then
    echo "✅ Snapshot build successful"
    echo "📁 Check the 'dist' directory for built binaries"
    ls -la dist/ 2>/dev/null || echo "No dist directory found"
else
    echo "❌ Snapshot build failed"
    exit 1
fi

echo "🎉 Goreleaser Docker configuration test completed successfully!"
echo ""
echo "To test Docker image building:"
echo "  docker build -t hello-gopher:test ."
echo ""
echo "To test full release (dry-run):"
echo "  goreleaser release --snapshot --clean"