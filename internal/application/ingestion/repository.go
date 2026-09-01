package ingestion

import (
	"context"

	"github.com/amirbakhtiari/godoc-ai/internal/domain"
)

type DocumentRepository interface {
	Create(ctx context.Context, document *domain.Document) error
}
