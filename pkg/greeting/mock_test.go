package greeting

import (
	"errors"
	"strings"
	"testing"
)

// TestMockGreeter demonstrates testability through interface mocking
func TestMockGreeter(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		setup    func(*MockGreeter)
	}{
		{
			name:     "default mock behavior",
			input:    "TestUser",
			expected: "Mock Hello, TestUser!",
			setup:    nil,
		},
		{
			name:     "custom mock behavior",
			input:    "CustomUser",
			expected: "Custom greeting for CustomUser",
			setup: func(m *MockGreeter) {
				m.GreetFunc = func(name string) string {
					return "Custom greeting for " + name
				}
			},
		},
		{
			name:     "empty name handling",
			input:    "",
			expected: "Mock Hello, MockGopher!",
			setup:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := NewMockGreeter()
			if tt.setup != nil {
				tt.setup(mock)
			}

			result := mock.Greet(tt.input)
			if result != tt.expected {
				t.Errorf("MockGreeter.Greet(%q) = %q, want %q", tt.input, result, tt.expected)
			}

			// Verify call logging
			expectedCall := `Greet("` + tt.input + `")`
			if len(mock.CallLog) != 1 || mock.CallLog[0] != expectedCall {
				t.Errorf("Expected call log [%q], got %v", expectedCall, mock.CallLog)
			}
		})
	}
}

// TestMockProverbProvider demonstrates proverb provider mocking
func TestMockProverbProvider(t *testing.T) {
	tests := []struct {
		name     string
		setup    func(*MockProverbProvider)
		expected string
	}{
		{
			name:     "default mock proverb",
			setup:    nil,
			expected: "Mock proverb: Don't communicate by sharing memory, share memory by communicating.",
		},
		{
			name: "custom proverbs",
			setup: func(m *MockProverbProvider) {
				m.SetProverbs([]string{"Custom test proverb"})
			},
			expected: "Custom test proverb",
		},
		{
			name: "empty proverbs",
			setup: func(m *MockProverbProvider) {
				m.SetProverbs([]string{})
			},
			expected: "No mock proverbs available",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := NewMockProverbProvider()
			if tt.setup != nil {
				tt.setup(mock)
			}

			result := mock.RandomProverb()
			if result != tt.expected {
				t.Errorf("MockProverbProvider.RandomProverb() = %q, want %q", result, tt.expected)
			}

			// Verify call logging
			if len(mock.CallLog) != 1 || mock.CallLog[0] != "RandomProverb()" {
				t.Errorf("Expected call log [RandomProverb()], got %v", mock.CallLog)
			}
		})
	}
}

// TestMockProverbProviderLoadProverbs tests the LoadProverbs mock functionality
func TestMockProverbProviderLoadProverbs(t *testing.T) {
	tests := []struct {
		name        string
		setup       func(*MockProverbProvider)
		expectError bool
	}{
		{
			name:        "successful load",
			setup:       nil,
			expectError: false,
		},
		{
			name: "load error",
			setup: func(m *MockProverbProvider) {
				m.LoadProverbsFunc = func() error {
					return errors.New("mock load error")
				}
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := NewMockProverbProvider()
			if tt.setup != nil {
				tt.setup(mock)
			}

			err := mock.LoadProverbs()
			if tt.expectError && err == nil {
				t.Error("Expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			// Verify call logging
			if len(mock.CallLog) != 1 || mock.CallLog[0] != "LoadProverbs()" {
				t.Errorf("Expected call log [LoadProverbs()], got %v", mock.CallLog)
			}
		})
	}
}

// TestMockService demonstrates combined interface mocking
func TestMockService(t *testing.T) {
	mock := NewMockService()

	// Test greeting functionality
	greeting := mock.Greet("TestUser")
	expectedGreeting := "Mock Hello, TestUser!"
	if greeting != expectedGreeting {
		t.Errorf("MockService.Greet() = %q, want %q", greeting, expectedGreeting)
	}

	// Test proverb functionality
	proverb := mock.RandomProverb()
	if !strings.Contains(proverb, "Mock proverb") {
		t.Errorf("MockService.RandomProverb() = %q, expected to contain 'Mock proverb'", proverb)
	}

	// Test load functionality
	err := mock.LoadProverbs()
	if err != nil {
		t.Errorf("MockService.LoadProverbs() unexpected error: %v", err)
	}

	// Verify both interfaces are implemented
	var _ Greeter = mock
	var _ ProverbProvider = mock
}

// TestErrorMockProverbProvider demonstrates error condition testing
func TestErrorMockProverbProvider(t *testing.T) {
	loadError := errors.New("failed to load proverbs")
	proverbError := "Error: No proverbs available"
	
	mock := NewErrorMockProverbProvider(loadError, proverbError)

	// Test error in LoadProverbs
	err := mock.LoadProverbs()
	if err != loadError {
		t.Errorf("ErrorMockProverbProvider.LoadProverbs() = %v, want %v", err, loadError)
	}

	// Test error in RandomProverb
	result := mock.RandomProverb()
	if result != proverbError {
		t.Errorf("ErrorMockProverbProvider.RandomProverb() = %q, want %q", result, proverbError)
	}

	// Verify call logging
	expectedCalls := []string{"LoadProverbs()", "RandomProverb()"}
	if len(mock.CallLog) != 2 {
		t.Errorf("Expected 2 calls, got %d", len(mock.CallLog))
	}
	for i, expectedCall := range expectedCalls {
		if i >= len(mock.CallLog) || mock.CallLog[i] != expectedCall {
			t.Errorf("Expected call %d to be %q, got %q", i, expectedCall, mock.CallLog[i])
		}
	}
}

// TestMockCallLogFunctionality tests the call logging features
func TestMockCallLogFunctionality(t *testing.T) {
	t.Run("greeter call log", func(t *testing.T) {
		mock := NewMockGreeter()
		
		// Make multiple calls
		mock.Greet("User1")
		mock.Greet("User2")
		mock.Greet("")
		
		expectedCalls := []string{
			`Greet("User1")`,
			`Greet("User2")`,
			`Greet("")`,
		}
		
		if len(mock.CallLog) != len(expectedCalls) {
			t.Errorf("Expected %d calls, got %d", len(expectedCalls), len(mock.CallLog))
		}
		
		for i, expected := range expectedCalls {
			if i >= len(mock.CallLog) || mock.CallLog[i] != expected {
				t.Errorf("Call %d: expected %q, got %q", i, expected, mock.CallLog[i])
			}
		}
		
		// Test clear functionality
		mock.ClearCallLog()
		if len(mock.CallLog) != 0 {
			t.Errorf("Expected empty call log after clear, got %v", mock.CallLog)
		}
	})
	
	t.Run("proverb provider call log", func(t *testing.T) {
		mock := NewMockProverbProvider()
		
		// Make multiple calls
		mock.LoadProverbs()
		mock.RandomProverb()
		mock.RandomProverb()
		
		expectedCalls := []string{
			"LoadProverbs()",
			"RandomProverb()",
			"RandomProverb()",
		}
		
		if len(mock.CallLog) != len(expectedCalls) {
			t.Errorf("Expected %d calls, got %d", len(expectedCalls), len(mock.CallLog))
		}
		
		for i, expected := range expectedCalls {
			if i >= len(mock.CallLog) || mock.CallLog[i] != expected {
				t.Errorf("Call %d: expected %q, got %q", i, expected, mock.CallLog[i])
			}
		}
		
		// Test clear functionality
		mock.ClearCallLog()
		if len(mock.CallLog) != 0 {
			t.Errorf("Expected empty call log after clear, got %v", mock.CallLog)
		}
	})
}

// TestInterfaceCompliance verifies that mocks implement the required interfaces
func TestInterfaceCompliance(t *testing.T) {
	// Test that MockGreeter implements Greeter
	var _ Greeter = (*MockGreeter)(nil)
	
	// Test that MockProverbProvider implements ProverbProvider
	var _ ProverbProvider = (*MockProverbProvider)(nil)
	
	// Test that ErrorMockProverbProvider implements ProverbProvider
	var _ ProverbProvider = (*ErrorMockProverbProvider)(nil)
	
	// Test that MockService implements both interfaces
	var _ Greeter = (*MockService)(nil)
	var _ ProverbProvider = (*MockService)(nil)
	
	// If we get here without compilation errors, the interfaces are properly implemented
	t.Log("All mock implementations properly implement their respective interfaces")
}