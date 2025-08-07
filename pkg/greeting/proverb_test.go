package greeting

import (
	"strings"
	"testing"
)

func TestLoadProverbs(t *testing.T) {
	tests := []struct {
		name        string
		expectError bool
		minProverbs int
	}{
		{
			name:        "load embedded proverbs successfully",
			expectError: false,
			minProverbs: 50, // We expect at least 50 proverbs as per requirements
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := NewService()
			err := service.LoadProverbs()

			if tt.expectError && err == nil {
				t.Errorf("LoadProverbs() expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("LoadProverbs() unexpected error: %v", err)
			}

			if !tt.expectError {
				if len(service.proverbs) < tt.minProverbs {
					t.Errorf("LoadProverbs() loaded %d proverbs, expected at least %d", len(service.proverbs), tt.minProverbs)
				}

				// Verify all proverbs are non-empty
				for i, proverb := range service.proverbs {
					if strings.TrimSpace(proverb) == "" {
						t.Errorf("LoadProverbs() proverb at index %d is empty", i)
					}
				}
			}
		})
	}
}

func TestRandomProverb(t *testing.T) {
	tests := []struct {
		name           string
		setupService   func() *Service
		expectContains string
		expectError    bool
	}{
		{
			name: "returns random proverb after loading",
			setupService: func() *Service {
				service := NewService()
				err := service.LoadProverbs()
				if err != nil {
					t.Fatalf("Failed to load proverbs: %v", err)
				}
				return service
			},
			expectContains: "", // Any non-empty string is valid
			expectError:    false,
		},
		{
			name: "auto-loads proverbs if not loaded",
			setupService: func() *Service {
				return NewService() // Don't pre-load proverbs
			},
			expectContains: "", // Any non-empty string is valid
			expectError:    false,
		},
		{
			name: "handles empty proverb list gracefully",
			setupService: func() *Service {
				service := NewService()
				service.proverbs = []string{} // Empty proverb list
				return service
			},
			expectContains: "", // Should auto-load
			expectError:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := tt.setupService()
			result := strings.TrimSpace(service.RandomProverb())

			if result == "" {
				t.Errorf("RandomProverb() returned empty string")
			}

			// Verify it doesn't return error messages for normal cases
			if !tt.expectError && strings.Contains(result, "Error") {
				t.Errorf("RandomProverb() returned error message: %s", result)
			}

			// Test that multiple calls can return different proverbs (probabilistic test)
			if len(service.proverbs) > 1 {
				results := make(map[string]bool)
				for i := 0; i < 10; i++ {
					proverb := service.RandomProverb()
					results[proverb] = true
				}
				// With 10 calls and multiple proverbs, we should get some variety
				// This is probabilistic, but with 60+ proverbs, it's very likely
				if len(results) == 1 && len(service.proverbs) > 5 {
					t.Logf("Warning: RandomProverb() returned same result 10 times (could be random chance)")
				}
			}
		})
	}
}

func TestRandomProverbConsistency(t *testing.T) {
	service := NewService()
	err := service.LoadProverbs()
	if err != nil {
		t.Fatalf("Failed to load proverbs: %v", err)
	}

	// Test that all returned proverbs are from the loaded set
	proverbSet := make(map[string]bool)
	for _, proverb := range service.proverbs {
		proverbSet[proverb] = true
	}

	for i := 0; i < 20; i++ {
		result := service.RandomProverb()
		if !proverbSet[result] {
			t.Errorf("RandomProverb() returned proverb not in loaded set: %s", result)
		}
	}
}

// TestProverbDataIntegrity verifies the embedded proverb data meets requirements
func TestProverbDataIntegrity(t *testing.T) {
	if proverbData == "" {
		t.Fatal("Embedded proverb data is empty")
	}

	lines := strings.Split(strings.TrimSpace(proverbData), "\n")
	validProverbs := 0

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" && !strings.HasPrefix(line, "#") {
			validProverbs++
		}
	}

	if validProverbs < 50 {
		t.Errorf("Embedded data contains %d valid proverbs, expected at least 50", validProverbs)
	}
}

// TestEmptyProverbDataHandling tests edge case with mock empty data
func TestEmptyProverbDataHandling(t *testing.T) {
	// Create a service and manually set empty proverbs to test error handling
	service := NewService()
	
	// Simulate the case where embedded data would be empty
	// We can't easily mock the embedded data, but we can test the error path
	// by directly testing the LoadProverbs logic with empty service.proverbs
	
	// First, load normally to ensure the method works
	err := service.LoadProverbs()
	if err != nil {
		t.Fatalf("LoadProverbs() failed with valid data: %v", err)
	}
	
	// Now test RandomProverb with empty proverbs slice
	service.proverbs = []string{}
	result := service.RandomProverb()
	
	// Should auto-load and return a valid proverb
	if result == "" || strings.Contains(result, "Error") {
		t.Errorf("RandomProverb() with empty proverbs should auto-load, got: %s", result)
	}
}