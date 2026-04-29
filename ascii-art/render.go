package main

func RenderLine(text string, banner map[rune][]string) []string {

	var result = make([]string, 8)
	for i := 0; i < 8; i++ {
		var line string
		for _, ch := range text {
			if ch < 32 || ch > 126 {
				ch = ' '
			}
			line += banner[ch][i]
		}
		result[i] = line
	}
	return result
}
