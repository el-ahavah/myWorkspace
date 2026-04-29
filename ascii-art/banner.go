package main

import (
	"fmt"
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")

	if len(lines) < 855 {
		return nil, fmt.Errorf("invalid banner format")
	}

	bannerMap := make(map[rune][]string)

	for ch := rune(32); ch <= 126; ch++ {
		start := (ch-32)*9 + 1
		end := start + 8

		if int(end) > len(lines) {
			return nil, fmt.Errorf("invalid banner format")
		}
		bannerMap[ch] = lines[start:end]
	}
	return bannerMap, nil
}
