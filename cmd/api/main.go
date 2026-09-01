package main

import (
	"fmt"
	"log"

	"github.com/amirbakhtiari/godoc-ai/internal/application/ingestion"
	"github.com/amirbakhtiari/godoc-ai/internal/infrastructure/document"
)

func main() {
	loader := document.NewMarkdownLoader()
	scanner := ingestion.NewScanner(loader)
	docs, err := scanner.Scan("docs/")
	if err != nil {
		log.Fatal(err)
	}
	for _, doc := range docs {
		fmt.Println("-----------------------------------")
		fmt.Println("ID: ", doc.ID)
		fmt.Println("Title: ", doc.Title)
		fmt.Println("Source:  ", doc.Source)
		fmt.Println("-----------------------------------")
	}
}
