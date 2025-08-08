package main

import (
	"os"
	"os/exec"
	"testing"
)

// TestMain tests the main function by running the compiled binary
func TestMain(t *testing.T) {
	// This test verifies that the main function can be called without panicking
	// We can't easily test main() directly since it calls os.Exit, so we test
	// the compiled binary instead
	
	if testing.Short() {
		t.Skip("Skipping main test in short mode")
	}
	
	// Build the binary for testing
	cmd := exec.Command("go", "build", "-o", "hello-gopher-test.exe", ".")
	cmd.Dir = "."
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Failed to build test binary: %v", err)
	}
	
	// Clean up the test binary
	defer func() {
		os.Remove("hello-gopher-test.exe")
	}()
	
	// Test that the binary runs without crashing
	testCmd := exec.Command("./hello-gopher-test.exe", "--help")
	testCmd.Dir = "."
	output, err := testCmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Binary execution failed: %v, output: %s", err, output)
	}
	
	// Verify that help output contains expected content
	outputStr := string(output)
	if len(outputStr) == 0 {
		t.Error("Expected non-empty help output")
	}
}

// TestMainIntegration tests the main function integration
func TestMainIntegration(t *testing.T) {
	// Test that main doesn't panic when imported
	// This is a basic smoke test for the main package
	
	// If we can import and run this test, main.go is syntactically correct
	// and doesn't have import issues
	t.Log("Main package imported successfully")
}

// TestMainFunctionExists tests that main function exists and can be referenced
func TestMainFunctionExists(t *testing.T) {
	// This test ensures the main function exists and is properly defined
	// We can't call main() directly due to os.Exit, but we can verify
	// that the function exists and the package compiles correctly
	
	// Test that the main package compiles and imports work correctly
	// The fact that this test runs means main.go is syntactically correct
	t.Log("Main function exists and package compiles correctly")
}

// BenchmarkMainExecution benchmarks the main execution path
func BenchmarkMainExecution(b *testing.B) {
	if testing.Short() {
		b.Skip("Skipping benchmark in short mode")
	}
	
	// Build the binary once for benchmarking
	cmd := exec.Command("go", "build", "-o", "hello-gopher-bench.exe", ".")
	cmd.Dir = "."
	err := cmd.Run()
	if err != nil {
		b.Fatalf("Failed to build benchmark binary: %v", err)
	}
	
	// Clean up the benchmark binary
	defer func() {
		os.Remove("hello-gopher-bench.exe")
	}()
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		testCmd := exec.Command("./hello-gopher-bench.exe", "greet", "--name", "BenchUser")
		testCmd.Dir = "."
		_, err := testCmd.CombinedOutput()
		if err != nil {
			b.Fatalf("Benchmark execution failed: %v", err)
		}
	}
}