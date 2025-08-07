package greeting

import (
	"fmt"
	"testing"
)

func TestService_Greet(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "default name when empty string",
			input:    "",
			expected: "Hello, Gopher!",
		},
		{
			name:     "custom name provided",
			input:    "Alice",
			expected: "Hello, Alice!",
		},
		{
			name:     "another custom name",
			input:    "Bob",
			expected: "Hello, Bob!",
		},
		{
			name:     "name with special characters",
			input:    "José",
			expected: "Hello, José!",
		},
		{
			name:     "name with spaces",
			input:    "John Doe",
			expected: "Hello, John Doe!",
		},
		{
			name:     "name with numbers",
			input:    "User123",
			expected: "Hello, User123!",
		},
		{
			name:     "single character name",
			input:    "A",
			expected: "Hello, A!",
		},
		{
			name:     "long name",
			input:    "VeryLongNameForTesting",
			expected: "Hello, VeryLongNameForTesting!",
		},
	}

	service := NewService()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := service.Greet(tt.input)
			if result != tt.expected {
				t.Errorf("Greet(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestNewService(t *testing.T) {
	service := NewService()
	if service == nil {
		t.Error("NewService() returned nil")
	}

	// Verify service implements required interfaces
	var _ Greeter = service
	var _ ProverbProvider = service
}

func TestService_LoadProverbs(t *testing.T) {
	service := NewService()
	err := service.LoadProverbs()
	if err != nil {
		t.Errorf("LoadProverbs() returned error: %v", err)
	}
}

func TestService_RandomProverb(t *testing.T) {
	service := NewService()
	proverb := service.RandomProverb()
	if proverb == "" {
		t.Error("RandomProverb() returned empty string")
	}
}

// Benchmark test for greeting function performance
func BenchmarkService_Greet(b *testing.B) {
	service := NewService()
	for i := 0; i < b.N; i++ {
		service.Greet("TestUser")
	}
}

// Example test for documentation purposes
func ExampleService_Greet() {
	service := NewService()
	greeting := service.Greet("World")
	fmt.Println(greeting)
	// Output: Hello, World!
}

func ExampleService_Greet_defaultName() {
	service := NewService()
	greeting := service.Greet("")
	fmt.Println(greeting)
	// Output: Hello, Gopher!
}