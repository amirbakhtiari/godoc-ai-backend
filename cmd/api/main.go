package main

import (
	"context"
	"fmt"
	"log"

	"github.com/amirbakhtiari/godoc-ai/internal/application/ingestion"
	"github.com/amirbakhtiari/godoc-ai/internal/infrastructure/database"
	"github.com/amirbakhtiari/godoc-ai/internal/infrastructure/documents"
)

func main() {
	ctx := context.Background()

	cfg := database.Config{
		Host:     "localhost",
		Port:     "5432",
		User:     "godoc",
		Password: "godoc",
		Name:     "godoc",
	}

	db, err := database.NewPostgresPool(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println("Connected to postgres")

	loader := documents.NewMarkdownLoader()
	scanner := ingestion.NewScanner(loader)
	//repository := documents.NewPostgresRepository(db)
	docs, err := scanner.Scan("docs/")
	if err != nil {
		log.Fatal(err)
	}

	chunker := documents.NewMarkdownChunker(50)

	for _, doc := range docs {
		chunks, err := chunker.Chunk(&doc)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("=======================================")
		fmt.Println("Document:", doc.Title)
		fmt.Println("Chunks:", len(chunks))
		fmt.Println("=======================================")

		for _, chunk := range chunks {
			fmt.Println()
			fmt.Println("Chunk =>", chunk.Position)
			fmt.Println("Section =>", chunk.Metadata.Section)
			fmt.Println("Heading Path =>", chunk.Metadata.HeadingPath)
			fmt.Println("Content Type =>", chunk.Metadata.ContentType)
			fmt.Println("Source =>", chunk.Metadata.Source)
			fmt.Println("Content:")
			fmt.Println(chunk.Content)
		}
	}
}
