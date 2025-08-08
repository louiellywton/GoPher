package greeting

import (
	"errors"
	"strings"
	"testing"
)

// BenchmarkMockGreeter benchmarks mock greeter performance
func BenchmarkMockGreeter(b *testing.B) {
	mock := NewMockGreeter()
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		_ = mock.Greet("BenchUser")
	}
}

// BenchmarkMockProverbProvider benchmarks mock proverb provider performance
func BenchmarkMockProverbProvider(b *testing.B) {
	mock := NewMockProverbProvider()
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		_ = mock.RandomProverb()
	}
}

// BenchmarkMockService benchmarks combined mock service performance
func BenchmarkMockService(b *testing.B) {
	mock := NewMockService()
	b.ResetTimer()
	
	b.Run("Greet", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = mock.Greet("BenchUser")
		}
	})
	
	b.Run("RandomProverb", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = mock.RandomProverb()
		}
	})
}

// BenchmarkMockCallLogging benchmarks the overhead of call logging
func BenchmarkMockCallLogging(b *testing.B) {
	mock := NewMockGreeter()
	
	b.Run("WithLogging", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = mock.Greet("User")
		}
	})
	
	b.Run("ClearLog", func(b *testing.B) {
		// Fill up the log first
		for i := 0; i < 100; i++ {
			mock.Greet("User")
		}
		
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			mock.ClearCallLog()
		}
	})
}

// BenchmarkErrorMockProverbProvider benchmarks error mock performance
func BenchmarkErrorMockProverbProvider(b *testing.B) {
	mock := NewErrorMockProverbProvider(
		errors.New("benchmark error"),
		"Error: benchmark proverb error",
	)
	
	b.Run("RandomProverb", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = mock.RandomProverb()
		}
	})
	
	b.Run("LoadProverbs", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = mock.LoadProverbs()
		}
	})
}

// BenchmarkStringOperations benchmarks string operations used in greeting
func BenchmarkStringOperations(b *testing.B) {
	names := []string{"", "Alice", "Bob", "VeryLongNameForBenchmarking", "José"}
	
	b.Run("StringConcatenation", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			name := names[i%len(names)]
			if name == "" {
				name = "Gopher"
			}
			_ = "Hello, " + name + "!"
		}
	})
	
	b.Run("StringFormatting", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			name := names[i%len(names)]
			if name == "" {
				name = "Gopher"
			}
			_ = strings.Join([]string{"Hello, ", name, "!"}, "")
		}
	})
}

// BenchmarkProverbSelection benchmarks proverb selection algorithms
func BenchmarkProverbSelection(b *testing.B) {
	proverbs := []string{
		"Don't communicate by sharing memory, share memory by communicating.",
		"Concurrency is not parallelism.",
		"Channels orchestrate; mutexes serialize.",
		"The bigger the interface, the weaker the abstraction.",
		"Make the zero value useful.",
		"interface{} says nothing.",
		"Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.",
		"A little copying is better than a little dependency.",
		"Syscall must always be guarded with build tags.",
		"Cgo must always be guarded with build tags.",
	}
	
	mock := NewMockProverbProvider()
	mock.SetProverbs(proverbs)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mock.RandomProverb()
	}
}

// BenchmarkInterfaceMethodCalls benchmarks interface method call overhead
func BenchmarkInterfaceMethodCalls(b *testing.B) {
	var greeter Greeter = NewMockGreeter()
	var provider ProverbProvider = NewMockProverbProvider()
	
	b.Run("GreeterInterface", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = greeter.Greet("BenchUser")
		}
	})
	
	b.Run("ProverbProviderInterface", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = provider.RandomProverb()
		}
	})
}

// BenchmarkMemoryAllocations benchmarks memory allocation patterns
func BenchmarkMemoryAllocations(b *testing.B) {
	b.Run("MockCreation", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = NewMockGreeter()
		}
	})
	
	b.Run("MockServiceCreation", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = NewMockService()
		}
	})
	
	b.Run("CallLogGrowth", func(b *testing.B) {
		mock := NewMockGreeter()
		b.ResetTimer()
		
		for i := 0; i < b.N; i++ {
			mock.Greet("User")
		}
	})
}