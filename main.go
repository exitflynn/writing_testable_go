package main

import (
	"fmt"

	"github.com/workshop/writing-testable-go/shortener"
)

func main() {
	url := "https://example.com/very/long/path"

	mapping, err := shortener.CreateMapping(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Shortened: %s -> %s\n", mapping.ShortCode, mapping.LongURL)
}
