package main

import "fmt"

func main() {
	banner, err := LoadBanner("standard.txt")
	if err != nil {
		fmt.Println("invalid banner file:", err)
		return
	}

	RenderLine("A", banner)

	for _, line := range RenderLine("joyce", banner) {
		fmt.Println(line)
	}
}
