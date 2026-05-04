package main

import (
	"strings"
)

func GenerateArt(input string, banner map[rune][]string) string {
	if input == "" {
		return ""
	}

	choice := strings.ReplaceAll(input, "\\n", "")

	_, err := ValidateInput(choice)
	if err != nil {
		return err.Error()
	}

	splitter := SplitInput(input)

	var build []string

	for _, ch := range splitter {

		if ch == "" {
			build = append(build, "")
			continue
		}

		build = append(build, RenderLine(ch, banner)...)
	}
	result := strings.Join(build, "\n")

	if choice != "" {
		result += "\n"
	}
	return result
}
