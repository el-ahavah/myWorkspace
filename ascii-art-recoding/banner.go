package main

import (
	"fmt"
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {
	banner, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file")
	}
	bannerlines := strings.Split(string(banner), "\n")

	if len(bannerlines) < 855 {
		return nil, fmt.Errorf("invalid banner length")
	}

	charbuild := map[rune][]string{}

	for ch := rune(32); ch <= 126; ch++ {
		start := (ch-32)*9 + 1
		end := start + 8

		if int(end) > len(bannerlines) {
			return nil, fmt.Errorf("out of bound")
		}

		charbuild[ch] = bannerlines[start:end]
	}
	return charbuild, err
}
