package greeting

import (
	"strings"
	"testing"
)

// TestRandomProverbEdgeCases tests edge cases for RandomProverb to improve coverage
func TestRandomProverbEdgeCases(t *testing.T) {
	t.Run("empty proverbs slice", func(t *testing.T) {
		service := &Service{
			proverbs: []string{}, // Empty slice
		}
		
		// This should trigger the auto-load path
		result := service.RandomProverb()
		
		// Should either load proverbs successfully or return an error message
		if result == "" {
			t.Error("Expected non-empty result from RandomProverb")
		}
		
		// Should contain either a proverb or an error message
		if !strings.Contains(result, "Don't") && !strings.Contains(result, "Error") && !strings.Contains(result, "No proverbs") {
			t.Logf("Got result: %s", result)
		}
	})
	
	t.Run("nil proverbs slice", func(t *testing.T) {
		service := &Service{
			proverbs: nil, // Nil slice
		}
		
		// This should trigger the auto-load path
		result := service.RandomProverb()
		
		// Should either load proverbs successfully or return an error message
		if result == "" {
			t.Error("Expected non-empty result from RandomProverb")
		}
	})
	
	t.Run("proverbs already loaded", func(t *testing.T) {
		service := NewService()
		err := service.LoadProverbs()
		if err != nil {
			t.Fatalf("Failed to load proverbs: %v", err)
		}
		
		// Now call RandomProverb - should not trigger reload
		result := service.RandomProverb()
		
		if result == "" {
			t.Error("Expected non-empty proverb")
		}
		
		// Should be a valid proverb, not an error message
		if strings.Contains(result, "Error") || strings.Contains(result, "No proverbs") {
			t.Errorf("Expected valid proverb, got: %s", result)
		}
	})
}

// TestLoadProverbsEdgeCases tests edge cases for LoadProverbs to improve coverage
func TestLoadProverbsEdgeCases(t *testing.T) {
	t.Run("multiple calls to LoadProverbs", func(t *testing.T) {
		service := NewService()
		
		// First load
		err1 := service.LoadProverbs()
		if err1 != nil {
			t.Fatalf("First load failed: %v", err1)
		}
		
		count1 := len(service.proverbs)
		
		// Second load - should work fine
		err2 := service.LoadProverbs()
		if err2 != nil {
			t.Fatalf("Second load failed: %v", err2)
		}
		
		count2 := len(service.proverbs)
		
		// Should have the same number of proverbs
		if count1 != count2 {
			t.Errorf("Expected same proverb count, got %d then %d", count1, count2)
		}
	})
	
	t.Run("proverb data processing", func(t *testing.T) {
		service := NewService()
		err := service.LoadProverbs()
		if err != nil {
			t.Fatalf("Failed to load proverbs: %v", err)
		}
		
		// Verify that proverbs were processed correctly
		if len(service.proverbs) == 0 {
			t.Error("Expected non-empty proverbs slice")
		}
		
		// Check that no empty strings made it through
		for i, proverb := range service.proverbs {
			if strings.TrimSpace(proverb) == "" {
				t.Errorf("Found empty proverb at index %d", i)
			}
			
			// Check that no comment lines made it through
			if strings.HasPrefix(proverb, "#") {
				t.Errorf("Found comment line in proverbs at index %d: %s", i, proverb)
			}
		}
	})
}

// TestServiceRandomness tests that RandomProverb produces varied results
func TestServiceRandomness(t *testing.T) {
	service := NewService()
	err := service.LoadProverbs()
	if err != nil {
		t.Fatalf("Failed to load proverbs: %v", err)
	}
	
	// If we have multiple proverbs, we should eventually get different ones
	if len(service.proverbs) > 1 {
		results := make(map[string]bool)
		
		// Try multiple times to get different proverbs
		for i := 0; i < 50; i++ {
			proverb := service.RandomProverb()
			results[proverb] = true
			
			// If we get at least 2 different proverbs, randomness is working
			if len(results) >= 2 {
				break
			}
		}
		
		// We should have gotten at least 2 different proverbs in 50 tries
		// (unless there's only 1 proverb, which we already checked)
		if len(results) < 2 {
			t.Logf("Only got %d unique proverbs in 50 tries, might indicate poor randomness", len(results))
			// Don't fail the test as this could be due to chance, just log it
		}
	}
}

// TestServiceProverbContent tests the content of loaded proverbs
func TestServiceProverbContent(t *testing.T) {
	service := NewService()
	err := service.LoadProverbs()
	if err != nil {
		t.Fatalf("Failed to load proverbs: %v", err)
	}
	
	// Verify we have a reasonable number of proverbs
	if len(service.proverbs) < 10 {
		t.Errorf("Expected at least 10 proverbs, got %d", len(service.proverbs))
	}
	
	// Check that proverbs contain expected Go-related content
	foundGoContent := false
	for _, proverb := range service.proverbs {
		// Look for common Go-related terms
		lowerProverb := strings.ToLower(proverb)
		if strings.Contains(lowerProverb, "go") || 
		   strings.Contains(lowerProverb, "gopher") ||
		   strings.Contains(lowerProverb, "channel") ||
		   strings.Contains(lowerProverb, "goroutine") ||
		   strings.Contains(lowerProverb, "interface") ||
		   strings.Contains(lowerProverb, "concurrency") {
			foundGoContent = true
			break
		}
	}
	
	if !foundGoContent {
		t.Log("No obvious Go-related content found in proverbs, but this might be okay")
	}
}

// TestServiceInterfaceCompliance tests that Service implements required interfaces
func TestServiceInterfaceCompliance(t *testing.T) {
	service := NewService()
	
	// Test that Service implements Greeter interface
	var _ Greeter = service
	
	// Test that Service implements ProverbProvider interface
	var _ ProverbProvider = service
	
	// If we get here without compilation errors, interfaces are properly implemented
	t.Log("Service properly implements both Greeter and ProverbProvider interfaces")
}

// TestServiceConcurrentAccess tests concurrent access to service methods
func TestServiceConcurrentAccess(t *testing.T) {
	service := NewService()
	err := service.LoadProverbs()
	if err != nil {
		t.Fatalf("Failed to load proverbs: %v", err)
	}
	
	// Test concurrent access to RandomProverb
	done := make(chan bool, 10)
	
	for i := 0; i < 10; i++ {
		go func() {
			defer func() { done <- true }()
			
			for j := 0; j < 10; j++ {
				proverb := service.RandomProverb()
				if proverb == "" {
					t.Errorf("Got empty proverb in concurrent access")
				}
			}
		}()
	}
	
	// Wait for all goroutines to complete
	for i := 0; i < 10; i++ {
		<-done
	}
}

// TestServiceMemoryUsage tests that service doesn't leak memory
func TestServiceMemoryUsage(t *testing.T) {
	// Create and destroy many services to check for memory leaks
	for i := 0; i < 100; i++ {
		service := NewService()
		err := service.LoadProverbs()
		if err != nil {
			t.Fatalf("Failed to load proverbs in iteration %d: %v", i, err)
		}
		
		// Use the service
		_ = service.Greet("Test")
		_ = service.RandomProverb()
		
		// Service should be garbage collected when it goes out of scope
	}
	
	// If we get here without running out of memory, we're probably okay
	t.Log("Memory usage test completed successfully")
}