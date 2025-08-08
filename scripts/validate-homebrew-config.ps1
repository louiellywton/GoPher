# Validation script for Homebrew tap configuration (PowerShell version)
# This script checks that all necessary components are in place

# Function to print colored output
function Write-Status {
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

function Write-Info {
    param([string]$Message)
    Write-Host "ℹ $Message" -ForegroundColor Blue
}

# Check if .goreleaser.yaml exists and has homebrew configuration
function Test-GoreleaserConfig {
    Write-Info "Checking Goreleaser configuration..."
    
    if (-not (Test-Path ".goreleaser.yaml")) {
        Write-Error ".goreleaser.yaml not found"
        return $false
    }
    
    Write-Status ".goreleaser.yaml exists"
    
    $content = Get-Content ".goreleaser.yaml" -Raw
    
    # Check for brews section
    if ($content -match "^brews:") {
        Write-Status "Homebrew brews section found"
    }
    else {
        Write-Error "No 'brews:' section found in .goreleaser.yaml"
        return $false
    }
    
    # Check for repository configuration
    if ($content -match "name: homebrew-tap") {
        Write-Status "Homebrew tap repository configured"
    }
    else {
        Write-Error "Homebrew tap repository not configured"
        return $false
    }
    
    # Check for required fields
    $requiredFields = @("homepage:", "description:", "install:", "test:")
    foreach ($field in $requiredFields) {
        if ($content -match [regex]::Escape($field)) {
            Write-Status "Field '$field' configured"
        }
        else {
            Write-Warning "Field '$field' not found (may be optional)"
        }
    }
    
    return $true
}

# Check if documentation exists
function Test-Documentation {
    Write-Info "Checking documentation..."
    
    $docs = @(
        "docs/homebrew-tap-setup.md",
        "docs/homebrew-tap-repository-setup.md"
    )
    
    foreach ($doc in $docs) {
        if (Test-Path $doc) {
            Write-Status "$doc exists"
        }
        else {
            Write-Warning "$doc not found"
        }
    }
}

# Check if test scripts exist
function Test-Scripts {
    Write-Info "Checking test scripts..."
    
    $scripts = @(
        "scripts/test-homebrew-tap.sh",
        "scripts/test-homebrew-tap.ps1",
        "scripts/validate-homebrew-config.sh",
        "scripts/validate-homebrew-config.ps1"
    )
    
    foreach ($script in $scripts) {
        if (Test-Path $script) {
            Write-Status "$script exists"
        }
        else {
            Write-Warning "$script not found"
        }
    }
}

# Validate YAML syntax (basic check)
function Test-YamlSyntax {
    Write-Info "Validating YAML syntax..."
    
    try {
        $content = Get-Content ".goreleaser.yaml" -Raw
        
        # Basic YAML validation - check for common issues
        $lines = $content -split "`n"
        $lineNumber = 0
        $hasErrors = $false
        
        foreach ($line in $lines) {
            $lineNumber++
            
            # Check for tabs (YAML should use spaces)
            if ($line -match "`t") {
                Write-Warning "Line $lineNumber contains tabs (YAML should use spaces only)"
            }
            
            # Check for basic structure issues
            if ($line -match "^\s*-\s*$") {
                Write-Warning "Line $lineNumber has empty list item"
            }
        }
        
        if (-not $hasErrors) {
            Write-Status "Basic YAML structure looks good"
        }
    }
    catch {
        Write-Error "Error reading .goreleaser.yaml: $($_.Exception.Message)"
        return $false
    }
    
    return $true
}

# Check GitHub repository settings (requires gh CLI)
function Test-GitHubSettings {
    Write-Info "Checking GitHub repository settings..."
    
    $ghCommand = Get-Command gh -ErrorAction SilentlyContinue
    if ($ghCommand) {
        try {
            # Check if we're in a git repository
            $gitDir = git rev-parse --git-dir 2>$null
            if ($gitDir) {
                # Get repository info
                $repoInfo = gh repo view --json name,owner,isPrivate 2>$null | ConvertFrom-Json
                
                if ($repoInfo) {
                    $repoName = $repoInfo.name
                    $owner = $repoInfo.owner.login
                    $isPrivate = $repoInfo.isPrivate
                    
                    Write-Status "Repository: $owner/$repoName"
                    
                    if (-not $isPrivate) {
                        Write-Status "Repository is public (required for Homebrew)"
                    }
                    else {
                        Write-Warning "Repository is private (Homebrew requires public repositories)"
                    }
                }
                else {
                    Write-Warning "Could not fetch repository information"
                }
            }
            else {
                Write-Warning "Not in a git repository"
            }
        }
        catch {
            Write-Warning "Error checking GitHub settings: $($_.Exception.Message)"
        }
    }
    else {
        Write-Warning "GitHub CLI (gh) not available, skipping GitHub checks"
    }
}

# Check if required environment requirements are documented
function Test-EnvironmentRequirements {
    Write-Info "Checking environment requirements..."
    
    # Check if README mentions installation methods
    if (Test-Path "README.md") {
        $readmeContent = Get-Content "README.md" -Raw
        if ($readmeContent -match "brew install") {
            Write-Status "README.md mentions Homebrew installation"
        }
        else {
            Write-Warning "README.md doesn't mention Homebrew installation"
        }
    }
    else {
        Write-Warning "README.md not found"
    }
}

# Check if the goreleaser configuration matches the expected repository structure
function Test-ConfigurationConsistency {
    Write-Info "Checking configuration consistency..."
    
    if (Test-Path ".goreleaser.yaml") {
        $content = Get-Content ".goreleaser.yaml" -Raw
        
        # Check if owner matches expected value
        if ($content -match "owner: louiellywton") {
            Write-Status "Repository owner is correctly configured"
        }
        else {
            Write-Warning "Repository owner may not be correctly configured"
        }
        
        # Check if the binary name matches project
        if ($content -match "binary: hello-gopher") {
            Write-Status "Binary name is correctly configured"
        }
        else {
            Write-Warning "Binary name may not be correctly configured"
        }
        
        # Check if main path is correct
        if ($content -match "main: ./cmd/hello-gopher") {
            Write-Status "Main package path is correctly configured"
        }
        else {
            Write-Warning "Main package path may not be correctly configured"
        }
    }
}

# Main validation function
function Main {
    Write-Host "=== Homebrew Tap Configuration Validation ===" -ForegroundColor Cyan
    Write-Host "This script validates the Homebrew tap setup for hello-gopher"
    Write-Host ""
    
    $exitCode = 0
    
    # Run all checks
    if (-not (Test-GoreleaserConfig)) {
        $exitCode = 1
    }
    Write-Host ""
    
    Test-Documentation
    Write-Host ""
    
    Test-Scripts
    Write-Host ""
    
    if (-not (Test-YamlSyntax)) {
        $exitCode = 1
    }
    Write-Host ""
    
    Test-GitHubSettings
    Write-Host ""
    
    Test-EnvironmentRequirements
    Write-Host ""
    
    Test-ConfigurationConsistency
    Write-Host ""
    
    # Summary
    if ($exitCode -eq 0) {
        Write-Status "All critical checks passed!"
        Write-Host ""
        Write-Host "Next steps:" -ForegroundColor Yellow
        Write-Host "1. Create the homebrew-tap repository on GitHub"
        Write-Host "2. Run the test scripts to verify functionality"
        Write-Host "3. Create a release to test the automatic formula generation"
        Write-Host ""
        Write-Host "To create the homebrew-tap repository, follow the guide in:"
        Write-Host "docs/homebrew-tap-repository-setup.md"
    }
    else {
        Write-Error "Some critical checks failed. Please address the issues above."
        Write-Host ""
        Write-Host "Common fixes:" -ForegroundColor Yellow
        Write-Host "1. Ensure .goreleaser.yaml has the correct homebrew configuration"
        Write-Host "2. Verify repository settings and permissions"
        Write-Host "3. Check YAML syntax for any errors"
    }
    
    exit $exitCode
}

# Run main function
Main