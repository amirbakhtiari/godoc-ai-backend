package ingestion

import "github.com/amirbakhtiari/godoc-ai/internal/domain"

type Chunker interface {
	Chunk(document *domain.Document) ([]domain.Chunk, error)
}
