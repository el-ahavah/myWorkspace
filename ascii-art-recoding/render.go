package main

func RenderLine(text string, banner map[rune][]string) []string {
	var result []string

	for i := 0; i < 8; i++ {
		var char string
		for _, ch := range text {
			char += banner[ch][i]
		}
		result = append(result, char)
	}
	return result
}
