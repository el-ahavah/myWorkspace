package main

import (
	"testing"
)

// helper: creates a minimal valid banner for tests
func makeValidBanner() map[rune][]string {
	banner := make(map[rune][]string)

	for ch := rune(32); ch <= 126; ch++ {
		lines := make([]string, 8)
		for i := 0; i < 8; i++ {
			lines[i] = string(ch)
		}
		banner[ch] = lines
	}

	return banner
}

func TestValidateAll_ValidInput(t *testing.T) {
	banner := makeValidBanner()
	err := ValidateAll("Hello World", banner)
	if err != nil {
		t.Errorf("expected valid input, got error: %v", err)
	}
}

func TestValidateAll_ValidInputWithNewlines(t *testing.T) {
	banner := makeValidBanner()
	err := ValidateAll("Hello\nWorld", banner)
	if err != nil {
		t.Errorf("expected valid input with newline, got error: %v", err)
	}
}

func TestValidateAll_InvalidCharacter(t *testing.T) {
	banner := makeValidBanner()
	err := ValidateAll("Hello🙂", banner)

	if err == nil {
		t.Errorf("expected error for invalid character, got nil")
	}
}

func TestValidateAll_EmptyInput(t *testing.T) {
	banner := makeValidBanner()
	err := ValidateAll("", banner)

	if err != nil {
		t.Errorf("expected empty input to be valid, got error: %v", err)
	}
}

func TestValidateAll_BannerMissingCharacter(t *testing.T) {
	banner := makeValidBanner()

	// remove one required character
	delete(banner, 'A')

	err := ValidateAll("ABC", banner)
	if err == nil {
		t.Errorf("expected error for missing banner character, got nil")
	}
}

func TestValidateAll_BannerWrongHeight(t *testing.T) {
	banner := makeValidBanner()

	// corrupt one character
	banner['B'] = []string{"1", "2", "3"} // invalid height

	err := ValidateAll("BBB", banner)
	if err == nil {
		t.Errorf("expected error for invalid banner height, got nil")
	}
}

func TestValidateAll_BannerWrongSize(t *testing.T) {
	banner := make(map[rune][]string)

	// only partial banner (not 95 chars)
	for ch := rune(32); ch < 50; ch++ {
		banner[ch] = make([]string, 8)
	}

	err := ValidateAll("Hello", banner)
	if err == nil {
		t.Errorf("expected error for invalid banner size, got nil")
	}
}

func TestValidateAll_BoundaryCharacters(t *testing.T) {
	banner := makeValidBanner()

	// lowest and highest valid ASCII
	err := ValidateAll(string(rune(32))+string(rune(126)), banner)
	if err != nil {
		t.Errorf("expected boundary chars to be valid, got error: %v", err)
	}
}

func TestValidateAll_InvalidLowASCII(t *testing.T) {
	banner := makeValidBanner()

	err := ValidateAll(string(rune(31)), banner)
	if err == nil {
		t.Errorf("expected error for low ASCII character, got nil")
	}
}

func TestValidateAll_InvalidHighASCII(t *testing.T) {
	banner := makeValidBanner()

	err := ValidateAll(string(rune(127)), banner)
	if err == nil {
		t.Errorf("expected error for high ASCII character, got nil")
	}
}
