// Package greeting provides functionality for generating friendly greetings
// and displaying Go programming proverbs.
//
// Example usage:
//   service := greeting.NewService()
//   fmt.Println(service.Greet("World"))
//   fmt.Println(service.RandomProverb())
package greeting

import "fmt"

// Greeter interface defines the contract for greeting functionality
type Greeter interface {
	Greet(name string) string
}

// ProverbProvider interface allows for easy mocking of proverb data,
// ensuring isolated unit tests without file system dependencies
type ProverbProvider interface {
	RandomProverb() string
	LoadProverbs() error
}

// Service implements both Greeter and ProverbProvider interfaces
type Service struct {
	proverbs []string
}

// NewService creates a new greeting service instance
func NewService() *Service {
	return &Service{}
}

// Greet returns a greeting message for the given name
func (s *Service) Greet(name string) string {
	if name == "" {
		name = "Gopher"
	}
	return fmt.Sprintf("Hello, %s!", name)
}

// RandomProverb and LoadProverbs implementations are in proverb.go