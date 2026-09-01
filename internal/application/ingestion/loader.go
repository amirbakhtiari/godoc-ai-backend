package ingestion

import "github.com/amirbakhtiari/godoc-ai/internal/domain"

type DocumentLoader interface {
	Load(path string) (document *domain.Document, err error)
}
