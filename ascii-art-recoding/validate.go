package main

import "fmt"

func ValidateInput(s string) (rune, error) {
	for _, ch := range s {
		if ch < rune(32) || ch > 126 {
			return ch, fmt.Errorf("invalid character")
		}
	}
	return 0, nil
}
