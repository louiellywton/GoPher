package greeting

import (
	"errors"
	"fmt"
	"strings"
)

// ExampleMockGreeter demonstrates how to use the mock greeter for testing
func ExampleMockGreeter() {
	mock := NewMockGreeter()
	
	// Use the mock with default behavior
	greeting := mock.Greet("Alice")
	fmt.Println(greeting)
	
	// Check call log
	fmt.Printf("Calls made: %d\n", len(mock.GetCallLog()))
	
	// Output:
	// Mock Hello, Alice!
	// Calls made: 1
}

// ExampleMockGreeter_customBehavior demonstrates custom mock behavior
func ExampleMockGreeter_customBehavior() {
	mock := NewMockGreeter()
	
	// Customize the mock behavior
	mock.GreetFunc = func(name string) string {
		return fmt.Sprintf("Custom greeting for %s", name)
	}
	
	greeting := mock.Greet("Bob")
	fmt.Println(greeting)
	
	// Output:
	// Custom greeting for Bob
}

// ExampleMockProverbProvider demonstrates mock proverb provider usage
func ExampleMockProverbProvider() {
	mock := NewMockProverbProvider()
	
	// Get a proverb
	proverb := mock.RandomProverb()
	fmt.Println(strings.Contains(proverb, "Mock proverb"))
	
	// Load proverbs (mock implementation)
	err := mock.LoadProverbs()
	fmt.Printf("Load error: %v\n", err)
	
	// Output:
	// true
	// Load error: <nil>
}

// ExampleMockProverbProvider_customProverbs demonstrates setting custom proverbs
func ExampleMockProverbProvider_customProverbs() {
	mock := NewMockProverbProvider()
	
	// Set custom proverbs
	customProverbs := []string{
		"Test proverb 1",
		"Test proverb 2",
	}
	mock.SetProverbs(customProverbs)
	
	proverb := mock.RandomProverb()
	fmt.Println(proverb)
	
	// Output:
	// Test proverb 1
}

// ExampleMockService demonstrates the combined mock service
func ExampleMockService() {
	mock := NewMockService()
	
	// Use both interfaces
	greeting := mock.Greet("World")
	proverb := mock.RandomProverb()
	
	fmt.Println(strings.Contains(greeting, "Mock Hello"))
	fmt.Println(strings.Contains(proverb, "Mock proverb"))
	
	// Output:
	// true
	// true
}

// ExampleErrorMockProverbProvider demonstrates error simulation
func ExampleErrorMockProverbProvider() {
	loadErr := errors.New("failed to load")
	proverbErr := "No proverbs available"
	
	mock := NewErrorMockProverbProvider(loadErr, proverbErr)
	
	// Simulate load error
	err := mock.LoadProverbs()
	fmt.Printf("Load failed: %v\n", err != nil)
	
	// Simulate proverb error
	result := mock.RandomProverb()
	fmt.Println(result)
	
	// Output:
	// Load failed: true
	// No proverbs available
}

// ExampleMockGreeter_callLogging demonstrates call logging functionality
func ExampleMockGreeter_callLogging() {
	mock := NewMockGreeter()
	
	// Make several calls
	mock.Greet("Alice")
	mock.Greet("Bob")
	mock.Greet("")
	
	// Check the call log
	calls := mock.GetCallLog()
	fmt.Printf("Total calls: %d\n", len(calls))
	fmt.Printf("First call: %s\n", calls[0])
	fmt.Printf("Last call: %s\n", calls[len(calls)-1])
	
	// Clear the log
	mock.ClearCallLog()
	fmt.Printf("After clear: %d\n", len(mock.GetCallLog()))
	
	// Output:
	// Total calls: 3
	// First call: Greet("Alice")
	// Last call: Greet("")
	// After clear: 0
}

// ExampleMockProverbProvider_emptyProverbs demonstrates handling empty proverbs
func ExampleMockProverbProvider_emptyProverbs() {
	mock := NewMockProverbProvider()
	
	// Set empty proverbs list
	mock.SetProverbs([]string{})
	
	result := mock.RandomProverb()
	fmt.Println(result)
	
	// Output:
	// No mock proverbs available
}

// ExampleNewMockGreeter demonstrates creating a new mock greeter
func ExampleNewMockGreeter() {
	mock := NewMockGreeter()
	
	// Verify it implements the interface
	var _ Greeter = mock
	
	// Test default behavior
	greeting := mock.Greet("")
	fmt.Println(strings.Contains(greeting, "MockGopher"))
	
	// Output:
	// true
}

// ExampleNewMockProverbProvider demonstrates creating a new mock proverb provider
func ExampleNewMockProverbProvider() {
	mock := NewMockProverbProvider()
	
	// Verify it implements the interface
	var _ ProverbProvider = mock
	
	// Test that it has default proverbs
	proverb := mock.RandomProverb()
	fmt.Println(len(proverb) > 0)
	
	// Output:
	// true
}

// ExampleNewMockService demonstrates creating a combined mock service
func ExampleNewMockService() {
	mock := NewMockService()
	
	// Verify it implements both interfaces
	var _ Greeter = mock
	var _ ProverbProvider = mock
	
	// Test both functionalities
	greeting := mock.Greet("Test")
	proverb := mock.RandomProverb()
	
	fmt.Println(len(greeting) > 0 && len(proverb) > 0)
	
	// Output:
	// true
}