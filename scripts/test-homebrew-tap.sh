#!/bin/bash

# Test script for Homebrew tap integration
# This script helps verify that the Homebrew tap works correctly

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check if Homebrew is installed
check_homebrew() {
    print_status "Checking if Homebrew is installed..."
    if ! command -v brew &> /dev/null; then
        print_error "Homebrew is not installed. Please install it first:"
        echo "  /bin/bash -c \"\$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)\""
        exit 1
    fi
    print_status "Homebrew is installed: $(brew --version | head -n1)"
}

# Check if the tap exists
check_tap() {
    print_status "Checking if tap is already added..."
    if brew tap | grep -q "louiellywton/tap"; then
        print_status "Tap is already added"
        return 0
    else
        print_warning "Tap is not added yet"
        return 1
    fi
}

# Add the tap
add_tap() {
    print_status "Adding the Homebrew tap..."
    if brew tap louiellywton/tap; then
        print_status "Successfully added tap"
    else
        print_error "Failed to add tap"
        exit 1
    fi
}

# Check if hello-gopher is installed
check_installation() {
    print_status "Checking if hello-gopher is installed..."
    if command -v hello-gopher &> /dev/null; then
        print_status "hello-gopher is installed: $(hello-gopher --version)"
        return 0
    else
        print_warning "hello-gopher is not installed"
        return 1
    fi
}

# Install hello-gopher
install_package() {
    print_status "Installing hello-gopher..."
    if brew install hello-gopher; then
        print_status "Successfully installed hello-gopher"
    else
        print_error "Failed to install hello-gopher"
        exit 1
    fi
}

# Test the installation
test_installation() {
    print_status "Testing hello-gopher installation..."
    
    # Test version command
    print_status "Testing version command..."
    if hello-gopher --version; then
        print_status "Version command works"
    else
        print_error "Version command failed"
        exit 1
    fi
    
    # Test greet command
    print_status "Testing greet command..."
    if hello-gopher greet --name "Homebrew Test"; then
        print_status "Greet command works"
    else
        print_error "Greet command failed"
        exit 1
    fi
    
    # Test proverb command
    print_status "Testing proverb command..."
    if hello-gopher proverb; then
        print_status "Proverb command works"
    else
        print_error "Proverb command failed"
        exit 1
    fi
    
    print_status "All tests passed!"
}

# Audit the formula
audit_formula() {
    print_status "Auditing the Homebrew formula..."
    if brew audit --strict louiellywton/tap/hello-gopher; then
        print_status "Formula audit passed"
    else
        print_warning "Formula audit found issues (this might be expected for development versions)"
    fi
}

# Clean up function
cleanup() {
    print_status "Cleaning up..."
    if command -v hello-gopher &> /dev/null; then
        print_status "Uninstalling hello-gopher..."
        brew uninstall hello-gopher || true
    fi
    
    if brew tap | grep -q "louiellywton/tap"; then
        print_status "Removing tap..."
        brew untap louiellywton/tap || true
    fi
    print_status "Cleanup complete"
}

# Main function
main() {
    echo "=== Homebrew Tap Test Script ==="
    echo "This script will test the hello-gopher Homebrew tap integration"
    echo
    
    # Parse command line arguments
    case "${1:-}" in
        "cleanup")
            cleanup
            exit 0
            ;;
        "test-only")
            if ! check_installation; then
                print_error "hello-gopher is not installed. Run without arguments to install and test."
                exit 1
            fi
            test_installation
            exit 0
            ;;
        "help"|"-h"|"--help")
            echo "Usage: $0 [command]"
            echo
            echo "Commands:"
            echo "  (no args)  - Full test: add tap, install, and test"
            echo "  test-only  - Test existing installation only"
            echo "  cleanup    - Remove installation and tap"
            echo "  help       - Show this help message"
            exit 0
            ;;
    esac
    
    # Check prerequisites
    check_homebrew
    
    # Add tap if not already added
    if ! check_tap; then
        add_tap
    fi
    
    # Install if not already installed
    if ! check_installation; then
        install_package
    fi
    
    # Test the installation
    test_installation
    
    # Audit the formula
    audit_formula
    
    echo
    print_status "Homebrew tap test completed successfully!"
    echo
    echo "To clean up, run: $0 cleanup"
}

# Run main function
main "$@"