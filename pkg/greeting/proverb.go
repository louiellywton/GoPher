package greeting

import (
	_ "embed"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

//go:embed proverb.txt
var proverbData string

// LoadProverbs loads proverbs from embedded data
func (s *Service) LoadProverbs() error {
	if proverbData == "" {
		return fmt.Errorf("embedded proverb data is empty")
	}

	// Split the embedded data into individual proverbs
	lines := strings.Split(strings.TrimSpace(proverbData), "\n")
	s.proverbs = make([]string, 0, len(lines))

	for _, line := range lines {
		line = strings.TrimSpace(line)
		// Skip empty lines and comments
		if line != "" && !strings.HasPrefix(line, "#") {
			s.proverbs = append(s.proverbs, line)
		}
	}

	if len(s.proverbs) == 0 {
		return fmt.Errorf("no valid proverbs found in embedded data")
	}

	return nil
}

// RandomProverb returns a random Go proverb
func (s *Service) RandomProverb() string {
	if len(s.proverbs) == 0 {
		// Try to load proverbs if not already loaded
		if err := s.LoadProverbs(); err != nil {
			return "Error loading proverbs: " + err.Error()
		}
	}

	if len(s.proverbs) == 0 {
		return "No proverbs available"
	}

	// Use current time as seed for randomness
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(s.proverbs))
	return s.proverbs[index]
}