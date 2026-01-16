package main

import (
	"fmt"

	"github.com/workshop/writing-testable-go/server"
	"github.com/workshop/writing-testable-go/shortener"
)

func main() {
	store := shortener.NewMemoryStore()
	srv := server.New(store)

	fmt.Println("Server running at http://localhost:8080")
	srv.Run(":8080")
}
