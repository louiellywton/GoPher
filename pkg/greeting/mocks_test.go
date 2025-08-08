package greeting

import (
	"fmt"
)

// MockGreeter is a mock implementation of the Greeter interface for testing
type MockGreeter struct {
	GreetFunc func(name string) string
	CallLog   []string
}

// NewMockGreeter creates a new mock greeter with default behavior
func NewMockGreeter() *MockGreeter {
	return &MockGreeter{
		GreetFunc: func(name string) string {
			if name == "" {
				name = "MockGopher"
			}
			return fmt.Sprintf("Mock Hello, %s!", name)
		},
		CallLog: make([]string, 0),
	}
}

// Greet implements the Greeter interface
func (m *MockGreeter) Greet(name string) string {
	m.CallLog = append(m.CallLog, fmt.Sprintf("Greet(%q)", name))
	return m.GreetFunc(name)
}

// MockProverbProvider is a mock implementation of the ProverbProvider interface for testing
type MockProverbProvider struct {
	RandomProverbFunc func() string
	LoadProverbsFunc  func() error
	CallLog           []string
	proverbs          []string
}

// NewMockProverbProvider creates a new mock proverb provider with default behavior
func NewMockProverbProvider() *MockProverbProvider {
	return &MockProverbProvider{
		proverbs: []string{
			"Mock proverb 1: Don't communicate by sharing memory, share memory by communicating.",
			"Mock proverb 2: Concurrency is not parallelism.",
			"Mock proverb 3: Channels orchestrate; mutexes serialize.",
		},
		RandomProverbFunc: func() string {
			return "Mock proverb: Don't communicate by sharing memory, share memory by communicating."
		},
		LoadProverbsFunc: func() error {
			return nil
		},
		CallLog: make([]string, 0),
	}
}

// RandomProverb implements the ProverbProvider interface
func (m *MockProverbProvider) RandomProverb() string {
	m.CallLog = append(m.CallLog, "RandomProverb()")
	return m.RandomProverbFunc()
}

// LoadProverbs implements the ProverbProvider interface
func (m *MockProverbProvider) LoadProverbs() error {
	m.CallLog = append(m.CallLog, "LoadProverbs()")
	return m.LoadProverbsFunc()
}

// SetProverbs allows setting custom proverbs for testing
func (m *MockProverbProvider) SetProverbs(proverbs []string) {
	m.proverbs = proverbs
	m.RandomProverbFunc = func() string {
		if len(m.proverbs) == 0 {
			return "No mock proverbs available"
		}
		return m.proverbs[0] // Return first proverb for predictable testing
	}
}

// GetCallLog returns the log of method calls for verification
func (m *MockGreeter) GetCallLog() []string {
	return m.CallLog
}

// GetCallLog returns the log of method calls for verification
func (m *MockProverbProvider) GetCallLog() []string {
	return m.CallLog
}

// ClearCallLog clears the call log
func (m *MockGreeter) ClearCallLog() {
	m.CallLog = make([]string, 0)
}

// ClearCallLog clears the call log
func (m *MockProverbProvider) ClearCallLog() {
	m.CallLog = make([]string, 0)
}

// MockService combines both interfaces for comprehensive testing
type MockService struct {
	*MockGreeter
	*MockProverbProvider
}

// NewMockService creates a new mock service that implements both interfaces
func NewMockService() *MockService {
	return &MockService{
		MockGreeter:        NewMockGreeter(),
		MockProverbProvider: NewMockProverbProvider(),
	}
}

// ErrorMockProverbProvider is a mock that simulates error conditions
type ErrorMockProverbProvider struct {
	LoadError     error
	ProverbError  string
	CallLog       []string
}

// NewErrorMockProverbProvider creates a mock that returns errors
func NewErrorMockProverbProvider(loadError error, proverbError string) *ErrorMockProverbProvider {
	return &ErrorMockProverbProvider{
		LoadError:    loadError,
		ProverbError: proverbError,
		CallLog:      make([]string, 0),
	}
}

// RandomProverb returns an error message
func (e *ErrorMockProverbProvider) RandomProverb() string {
	e.CallLog = append(e.CallLog, "RandomProverb()")
	if e.ProverbError != "" {
		return e.ProverbError
	}
	return "Error: Mock error condition"
}

// LoadProverbs returns the configured error
func (e *ErrorMockProverbProvider) LoadProverbs() error {
	e.CallLog = append(e.CallLog, "LoadProverbs()")
	return e.LoadError
}

// GetCallLog returns the log of method calls for verification
func (e *ErrorMockProverbProvider) GetCallLog() []string {
	return e.CallLog
}