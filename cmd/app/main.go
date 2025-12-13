package main

import (
	"fmt"
	"log"

	"github.com/atrakic/go-project-template/internal/uuid"
	"github.com/atrakic/go-project-template/pkg/greeting"
)

func main() {
	fmt.Println(greeting.Hello())

	id, err := uuid.Generate()
	if err != nil {
		log.Fatalf("Failed to generate UUID: %v", err)
	}

	fmt.Printf("Generated UUID: %s\n", id)
}
