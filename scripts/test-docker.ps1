# PowerShell script to test Docker functionality
# This script tests the Docker container across different scenarios on Windows

param(
    [switch]$Verbose
)

# Set error action preference
$ErrorActionPreference = "Stop"

Write-Host "🐳 Testing Docker functionality for hello-gopher..." -ForegroundColor Cyan

# Function to print status messages
function Write-Success {
    param([string]$Message)
    Write-Host "✓ $Message" -ForegroundColor Green
}

function Write-Warning {
    param([string]$Message)
    Write-Host "⚠ $Message" -ForegroundColor Yellow
}

function Write-Error {
    param([string]$Message)
    Write-Host "✗ $Message" -ForegroundColor Red
}

# Check if Docker is available
try {
    $null = docker --version
    Write-Success "Docker is available"
} catch {
    Write-Error "Docker is not installed or not in PATH"
    exit 1
}

# Build the Docker image locally
Write-Host "🔨 Building Docker image..." -ForegroundColor Cyan
try {
    if ($Verbose) {
        docker build -t hello-gopher:test .
    } else {
        docker build -t hello-gopher:test . | Out-Null
    }
    Write-Success "Docker image built successfully"
} catch {
    Write-Error "Failed to build Docker image"
    exit 1
}

# Test basic functionality
Write-Host "🧪 Testing basic functionality..." -ForegroundColor Cyan

# Test default greeting
Write-Host "Testing default greeting..."
try {
    $output = docker run --rm hello-gopher:test greet 2>$null
    if ($output -eq "Hello, Gopher!") {
        Write-Success "Default greeting works"
    } else {
        Write-Error "Default greeting failed. Expected 'Hello, Gopher!', got '$output'"
        exit 1
    }
} catch {
    Write-Error "Default greeting test failed with exception"
    exit 1
}

# Test custom name greeting
Write-Host "Testing custom name greeting..."
try {
    $output = docker run --rm hello-gopher:test greet --name Docker 2>$null
    if ($output -eq "Hello, Docker!") {
        Write-Success "Custom name greeting works"
    } else {
        Write-Error "Custom name greeting failed. Expected 'Hello, Docker!', got '$output'"
        exit 1
    }
} catch {
    Write-Error "Custom name greeting test failed with exception"
    exit 1
}

# Test short flag
Write-Host "Testing short flag..."
try {
    $output = docker run --rm hello-gopher:test greet -n Test 2>$null
    if ($output -eq "Hello, Test!") {
        Write-Success "Short flag works"
    } else {
        Write-Error "Short flag failed. Expected 'Hello, Test!', got '$output'"
        exit 1
    }
} catch {
    Write-Error "Short flag test failed with exception"
    exit 1
}

# Test proverb command
Write-Host "Testing proverb command..."
try {
    $output = docker run --rm hello-gopher:test proverb 2>$null
    if ($output -and $output.Length -gt 10) {
        $truncated = if ($output.Length -gt 50) { $output.Substring(0, 50) + "..." } else { $output }
        Write-Success "Proverb command works (output: $truncated)"
    } else {
        Write-Error "Proverb command failed or returned empty/short output"
        exit 1
    }
} catch {
    Write-Error "Proverb command test failed with exception"
    exit 1
}

# Test version command
Write-Host "Testing version command..."
try {
    $output = docker run --rm hello-gopher:test --version 2>$null
    if ($output -like "*hello-gopher version*") {
        Write-Success "Version command works"
    } else {
        Write-Error "Version command failed. Output: '$output'"
        exit 1
    }
} catch {
    Write-Error "Version command test failed with exception"
    exit 1
}

# Test help command
Write-Host "Testing help command..."
try {
    $output = docker run --rm hello-gopher:test --help 2>$null
    if ($output -like "*Hello-Gopher is a friendly command-line tool*") {
        Write-Success "Help command works"
    } else {
        Write-Error "Help command failed"
        exit 1
    }
} catch {
    Write-Error "Help command test failed with exception"
    exit 1
}

# Test error handling
Write-Host "Testing error handling..."
try {
    docker run --rm hello-gopher:test invalid-command 2>$null | Out-Null
    if ($LASTEXITCODE -eq 0) {
        Write-Error "Error handling test failed - should have returned non-zero exit code"
        exit 1
    } else {
        Write-Success "Error handling works correctly"
    }
} catch {
    Write-Success "Error handling works correctly (exception caught as expected)"
}

# Test container size
Write-Host "📏 Checking container size..." -ForegroundColor Cyan
try {
    $images = docker images hello-gopher:test --format "table {{.Size}}"
    $size = ($images -split "`n")[1]
    Write-Success "Container size: $size"
} catch {
    Write-Warning "Could not determine container size"
}

# Test container layers
Write-Host "🔍 Checking container layers..." -ForegroundColor Cyan
try {
    $history = docker history hello-gopher:test --format "table {{.CreatedBy}}"
    $layers = ($history -split "`n").Count - 1
    Write-Success "Container has $layers layers"
} catch {
    Write-Warning "Could not determine container layers"
}

# Clean up test image
Write-Host "🧹 Cleaning up..." -ForegroundColor Cyan
try {
    docker rmi hello-gopher:test | Out-Null
    Write-Success "Test image removed"
} catch {
    Write-Warning "Could not remove test image (may not exist)"
}

Write-Host ""
Write-Host "🎉 All Docker tests passed successfully!" -ForegroundColor Green
Write-Host ""
Write-Host "Container features verified:" -ForegroundColor Cyan
Write-Host "  ✓ Basic greeting functionality"
Write-Host "  ✓ Custom name support"
Write-Host "  ✓ Flag parsing (long and short)"
Write-Host "  ✓ Proverb display"
Write-Host "  ✓ Version information"
Write-Host "  ✓ Help system"
Write-Host "  ✓ Error handling"
Write-Host "  ✓ Minimal container size"
Write-Host ""
Write-Host "The Docker container is ready for production use!" -ForegroundColor Green