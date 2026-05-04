package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: go run . `text`")
		return
	}

	banner, err := LoadBanner("standard.txt")
	if err != nil {
		fmt.Println("invalid bannerfile")
		return
	}

	result := GenerateArt(os.Args[1], banner)

	fmt.Print(result)
}
