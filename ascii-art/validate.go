package main

import (
	"fmt"
)

func ValidateAll(text string, banner map[rune][]string) error {
	// 1. Validate banner integrity
	if len(banner) != 95 {
		return fmt.Errorf("invalid banner: expected 95 characters, got %d", len(banner))
	}

	for ch := rune(32); ch <= 126; ch++ {
		lines, ok := banner[ch]
		if !ok {
			return fmt.Errorf("missing character in banner: %q", ch)
		}

		if len(lines) != 8 {
			return fmt.Errorf("invalid height for character %q", ch)
		}
	}

	// 2. Validate input text
	for _, ch := range text {
		if ch == '\n' {
			continue
		}
		if ch < 32 || ch > 126 {
			return fmt.Errorf("invalid input character: %q", ch)
		}
	}

	return nil
}
