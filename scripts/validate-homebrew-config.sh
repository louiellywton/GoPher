#!/bin/bash

# Validation script for Homebrew tap configuration
# This script checks that all necessary components are in place

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    echo -e "${GREEN}✓${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}⚠${NC} $1"
}

print_error() {
    echo -e "${RED}✗${NC} $1"
}

print_info() {
    echo -e "${BLUE}ℹ${NC} $1"
}

# Check if .goreleaser.yaml exists and has homebrew configuration
check_goreleaser_config() {
    print_info "Checking Goreleaser configuration..."
    
    if [[ ! -f ".goreleaser.yaml" ]]; then
        print_error ".goreleaser.yaml not found"
        return 1
    fi
    
    print_status ".goreleaser.yaml exists"
    
    # Check for brews section
    if grep -q "^brews:" .goreleaser.yaml; then
        print_status "Homebrew brews section found"
    else
        print_error "No 'brews:' section found in .goreleaser.yaml"
        return 1
    fi
    
    # Check for repository configuration
    if grep -q "name: homebrew-tap" .goreleaser.yaml; then
        print_status "Homebrew tap repository configured"
    else
        print_error "Homebrew tap repository not configured"
        return 1
    fi
    
    # Check for required fields
    local required_fields=("homepage:" "description:" "install:" "test:")
    for field in "${required_fields[@]}"; do
        if grep -q "$field" .goreleaser.yaml; then
            print_status "Field '$field' configured"
        else
            print_warning "Field '$field' not found (may be optional)"
        fi
    done
    
    return 0
}

# Check if documentation exists
check_documentation() {
    print_info "Checking documentation..."
    
    local docs=(
        "docs/homebrew-tap-setup.md"
        "docs/homebrew-tap-repository-setup.md"
    )
    
    for doc in "${docs[@]}"; do
        if [[ -f "$doc" ]]; then
            print_status "$doc exists"
        else
            print_warning "$doc not found"
        fi
    done
}

# Check if test scripts exist
check_test_scripts() {
    print_info "Checking test scripts..."
    
    local scripts=(
        "scripts/test-homebrew-tap.sh"
        "scripts/test-homebrew-tap.ps1"
    )
    
    for script in "${scripts[@]}"; do
        if [[ -f "$script" ]]; then
            print_status "$script exists"
            if [[ -x "$script" ]]; then
                print_status "$script is executable"
            else
                print_warning "$script is not executable (run: chmod +x $script)"
            fi
        else
            print_warning "$script not found"
        fi
    done
}

# Validate YAML syntax (if yq is available)
validate_yaml_syntax() {
    print_info "Validating YAML syntax..."
    
    if command -v yq &> /dev/null; then
        if yq eval '.goreleaser.yaml' .goreleaser.yaml > /dev/null 2>&1; then
            print_status "YAML syntax is valid"
        else
            print_error "YAML syntax error in .goreleaser.yaml"
            return 1
        fi
    else
        print_warning "yq not available, skipping YAML validation"
    fi
}

# Check GitHub repository settings (requires gh CLI)
check_github_settings() {
    print_info "Checking GitHub repository settings..."
    
    if command -v gh &> /dev/null; then
        # Check if we're in a git repository
        if git rev-parse --git-dir > /dev/null 2>&1; then
            # Get repository info
            local repo_info
            repo_info=$(gh repo view --json name,owner,isPrivate 2>/dev/null || echo "")
            
            if [[ -n "$repo_info" ]]; then
                local repo_name
                repo_name=$(echo "$repo_info" | jq -r '.name' 2>/dev/null || echo "unknown")
                local owner
                owner=$(echo "$repo_info" | jq -r '.owner.login' 2>/dev/null || echo "unknown")
                local is_private
                is_private=$(echo "$repo_info" | jq -r '.isPrivate' 2>/dev/null || echo "unknown")
                
                print_status "Repository: $owner/$repo_name"
                
                if [[ "$is_private" == "false" ]]; then
                    print_status "Repository is public (required for Homebrew)"
                else
                    print_warning "Repository is private (Homebrew requires public repositories)"
                fi
            else
                print_warning "Could not fetch repository information"
            fi
        else
            print_warning "Not in a git repository"
        fi
    else
        print_warning "GitHub CLI (gh) not available, skipping GitHub checks"
    fi
}

# Check if required environment variables are documented
check_environment_requirements() {
    print_info "Checking environment requirements..."
    
    # Check if README mentions installation methods
    if [[ -f "README.md" ]]; then
        if grep -q "brew install" README.md; then
            print_status "README.md mentions Homebrew installation"
        else
            print_warning "README.md doesn't mention Homebrew installation"
        fi
    else
        print_warning "README.md not found"
    fi
}

# Main validation function
main() {
    echo "=== Homebrew Tap Configuration Validation ==="
    echo "This script validates the Homebrew tap setup for hello-gopher"
    echo
    
    local exit_code=0
    
    # Run all checks
    check_goreleaser_config || exit_code=1
    echo
    
    check_documentation
    echo
    
    check_test_scripts
    echo
    
    validate_yaml_syntax || exit_code=1
    echo
    
    check_github_settings
    echo
    
    check_environment_requirements
    echo
    
    # Summary
    if [[ $exit_code -eq 0 ]]; then
        print_status "All critical checks passed!"
        echo
        echo "Next steps:"
        echo "1. Create the homebrew-tap repository on GitHub"
        echo "2. Run the test scripts to verify functionality"
        echo "3. Create a release to test the automatic formula generation"
    else
        print_error "Some critical checks failed. Please address the issues above."
        echo
        echo "Common fixes:"
        echo "1. Ensure .goreleaser.yaml has the correct homebrew configuration"
        echo "2. Verify repository settings and permissions"
        echo "3. Check YAML syntax for any errors"
    fi
    
    exit $exit_code
}

# Run main function
main "$@"