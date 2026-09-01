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
		Host: "localhost",
		Port: "5432",
		User: "godoc",
		Pass: "godoc",
		Name: "godoc",
	}

	db, err := database.NewPostgresPool(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println("Connected to postgres")

	loader := documents.NewMarkdownLoader()
	scanner := ingestion.NewScanner(loader)
	repository := documents.NewPostgresRepository(db)

	docs, err := scanner.Scan("docs/")
	if err != nil {
		log.Fatal(err)
	}
	for _, doc := range docs {
		if err := repository.Create(ctx, &doc); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Inserted => ", doc.Title)
	}
}
