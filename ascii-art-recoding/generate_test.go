package main

import (
	"strings"
	"testing"
)

func TestGenerateArt(t *testing.T) {
	// Setup: Load the banner once for all tests.
	banner, err := LoadBanner("standard.txt")
	if err != nil {
		t.Fatalf("Setup failed: could not load standard.txt: %v", err)
	}

	// Define test scenarios based on project requirements.
	tests := []struct {
		name     string
		input    string
		validate func(t *testing.T, output string)
	}{
		{
			name:  "Empty Input produces nothing",
			input: "",
			validate: func(t *testing.T, output string) {
				if output != "" {
					t.Errorf("Expected empty string, got %q", output)
				}
			},
		},
		{
			name:  "Single newline literal produces one blank line",
			input: `\n`,
			validate: func(t *testing.T, output string) {
				if output != "\n" {
					t.Errorf("Expected exactly one newline, got %q", output)
				}
			},
		},
		{
			name:  "Multiple newline literals produce multiple blank lines",
			input: `\n\n\n`,
			validate: func(t *testing.T, output string) {
				if output != "\n\n\n" {
					t.Errorf("Expected three newlines, got %d", len(output))
				}
			},
		},
		{
			name:  "Mixed text and newlines 'A\\nB'",
			input: `A\nB`,
			validate: func(t *testing.T, output string) {
				// Should have 16 lines total (8 for A, 8 for B)
				lines := strings.Split(output, "\n")
				// strings.Split adds an extra empty element if the string ends in \n
				if len(lines) != 17 {
					t.Errorf("Expected 16 lines of art, found %d", len(lines)-1)
				}
			},
		},
		{
			name:  "Intermediate newline gap 'A\\n\\nB'",
			input: `A\n\nB`,
			validate: func(t *testing.T, output string) {
				// Should have 8 lines (A) + 1 blank line + 8 lines (B) = 17 total printed lines
				lines := strings.Split(output, "\n")
				if len(lines) != 18 {
					t.Errorf("Expected 17 lines total, got %d", len(lines)-1)
				}
				// The 9th line (index 8) should be the empty gap
				if lines[8] != "" {
					t.Errorf("Expected a blank line gap at index 8, got %q", lines[8])
				}
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := GenerateArt(tc.input, banner)
			tc.validate(t, result)
		})
	}
}
