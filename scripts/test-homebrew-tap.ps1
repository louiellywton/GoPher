# Test script for Homebrew tap integration (PowerShell version)
# This script helps verify that the Homebrew tap works correctly on Windows

param(
    [string]$Command = ""
)

# Function to print colored output
function Write-Status {
    param([string]$Message)
    Write-Host "[INFO] $Message" -ForegroundColor Green
}

function Write-Warning {
    param([string]$Message)
    Write-Host "[WARN] $Message" -ForegroundColor Yellow
}

function Write-Error {
    param([string]$Message)
    Write-Host "[ERROR] $Message" -ForegroundColor Red
}

# Check if running on Windows with WSL or if Homebrew is available
function Test-Homebrew {
    Write-Status "Checking if Homebrew is available..."
    
    # Check if we're in WSL
    if ($env:WSL_DISTRO_NAME) {
        Write-Status "Running in WSL environment"
        $brewPath = Get-Command brew -ErrorAction SilentlyContinue
        if ($brewPath) {
            Write-Status "Homebrew is available: $(brew --version | Select-Object -First 1)"
            return $true
        }
    }
    
    # Check if Homebrew is installed on Windows (via WSL or other means)
    try {
        $version = & brew --version 2>$null | Select-Object -First 1
        if ($version) {
            Write-Status "Homebrew is available: $version"
            return $true
        }
    }
    catch {
        # Homebrew not found
    }
    
    Write-Warning "Homebrew is not available on this Windows system."
    Write-Host "To test the Homebrew tap, you can:"
    Write-Host "1. Use WSL (Windows Subsystem for Linux) with Homebrew installed"
    Write-Host "2. Use a macOS or Linux system"
    Write-Host "3. Test the formula syntax online at: https://github.com/louiellywton/homebrew-tap"
    return $false
}

# Check if the tap exists
function Test-Tap {
    Write-Status "Checking if tap is already added..."
    try {
        $taps = & brew tap 2>$null
        if ($taps -match "louiellywton/tap") {
            Write-Status "Tap is already added"
            return $true
        }
        else {
            Write-Warning "Tap is not added yet"
            return $false
        }
    }
    catch {
        Write-Error "Failed to check taps"
        return $false
    }
}

# Add the tap
function Add-Tap {
    Write-Status "Adding the Homebrew tap..."
    try {
        & brew tap louiellywton/tap
        Write-Status "Successfully added tap"
        return $true
    }
    catch {
        Write-Error "Failed to add tap"
        return $false
    }
}

# Check if hello-gopher is installed
function Test-Installation {
    Write-Status "Checking if hello-gopher is installed..."
    try {
        $version = & hello-gopher --version 2>$null
        if ($version) {
            Write-Status "hello-gopher is installed: $version"
            return $true
        }
    }
    catch {
        # Not installed
    }
    
    Write-Warning "hello-gopher is not installed"
    return $false
}

# Install hello-gopher
function Install-Package {
    Write-Status "Installing hello-gopher..."
    try {
        & brew install hello-gopher
        Write-Status "Successfully installed hello-gopher"
        return $true
    }
    catch {
        Write-Error "Failed to install hello-gopher"
        return $false
    }
}

# Test the installation
function Test-HelloGopher {
    Write-Status "Testing hello-gopher installation..."
    
    # Test version command
    Write-Status "Testing version command..."
    try {
        & hello-gopher --version
        Write-Status "Version command works"
    }
    catch {
        Write-Error "Version command failed"
        return $false
    }
    
    # Test greet command
    Write-Status "Testing greet command..."
    try {
        & hello-gopher greet --name "Homebrew Test"
        Write-Status "Greet command works"
    }
    catch {
        Write-Error "Greet command failed"
        return $false
    }
    
    # Test proverb command
    Write-Status "Testing proverb command..."
    try {
        & hello-gopher proverb
        Write-Status "Proverb command works"
    }
    catch {
        Write-Error "Proverb command failed"
        return $false
    }
    
    Write-Status "All tests passed!"
    return $true
}

# Audit the formula
function Test-Formula {
    Write-Status "Auditing the Homebrew formula..."
    try {
        & brew audit --strict louiellywton/tap/hello-gopher
        Write-Status "Formula audit passed"
    }
    catch {
        Write-Warning "Formula audit found issues (this might be expected for development versions)"
    }
}

# Clean up function
function Remove-Installation {
    Write-Status "Cleaning up..."
    
    if (Test-Installation) {
        Write-Status "Uninstalling hello-gopher..."
        try {
            & brew uninstall hello-gopher
        }
        catch {
            Write-Warning "Failed to uninstall hello-gopher"
        }
    }
    
    if (Test-Tap) {
        Write-Status "Removing tap..."
        try {
            & brew untap louiellywton/tap
        }
        catch {
            Write-Warning "Failed to remove tap"
        }
    }
    
    Write-Status "Cleanup complete"
}

# Show help
function Show-Help {
    Write-Host "Homebrew Tap Test Script (PowerShell)"
    Write-Host "Usage: .\test-homebrew-tap.ps1 [command]"
    Write-Host ""
    Write-Host "Commands:"
    Write-Host "  (no args)  - Full test: add tap, install, and test"
    Write-Host "  test-only  - Test existing installation only"
    Write-Host "  cleanup    - Remove installation and tap"
    Write-Host "  help       - Show this help message"
    Write-Host ""
    Write-Host "Note: This script requires Homebrew to be available on Windows"
    Write-Host "      (typically through WSL or other Linux compatibility layer)"
}

# Main function
function Main {
    Write-Host "=== Homebrew Tap Test Script (PowerShell) ===" -ForegroundColor Cyan
    Write-Host "This script will test the hello-gopher Homebrew tap integration"
    Write-Host ""
    
    # Parse command line arguments
    switch ($Command.ToLower()) {
        "cleanup" {
            if (Test-Homebrew) {
                Remove-Installation
            }
            return
        }
        "test-only" {
            if (-not (Test-Homebrew)) {
                return
            }
            if (-not (Test-Installation)) {
                Write-Error "hello-gopher is not installed. Run without arguments to install and test."
                return
            }
            Test-HelloGopher
            return
        }
        "help" {
            Show-Help
            return
        }
        "" {
            # Continue with full test
        }
        default {
            Write-Error "Unknown command: $Command"
            Show-Help
            return
        }
    }
    
    # Check prerequisites
    if (-not (Test-Homebrew)) {
        return
    }
    
    # Add tap if not already added
    if (-not (Test-Tap)) {
        if (-not (Add-Tap)) {
            return
        }
    }
    
    # Install if not already installed
    if (-not (Test-Installation)) {
        if (-not (Install-Package)) {
            return
        }
    }
    
    # Test the installation
    if (-not (Test-HelloGopher)) {
        return
    }
    
    # Audit the formula
    Test-Formula
    
    Write-Host ""
    Write-Status "Homebrew tap test completed successfully!"
    Write-Host ""
    Write-Host "To clean up, run: .\test-homebrew-tap.ps1 cleanup"
}

# Run main function
Main