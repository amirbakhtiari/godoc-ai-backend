package documents

import (
	"fmt"
	"strings"

	"github.com/amirbakhtiari/godoc-ai/internal/domain"
)

type simpleChunker struct {
	maxWords int
}

func NewSimpleChunker(maxWords int) simpleChunker {
	return simpleChunker{
		maxWords: maxWords,
	}
}
func (c simpleChunker) Chunk(document *domain.Document) ([]domain.Chunk, error) {
	words := strings.Fields(document.Content)

	if len(words) == 0 {
		return nil, nil
	}

	var chunks []domain.Chunk
	for start := 0; start < len(words); start += c.maxWords {
		end := start + c.maxWords
		if end > len(words) {
			end = len(words)
		}

		content := strings.Join(words[start:end], " ")

		chunk := domain.Chunk{
			ID:         fmt.Sprintf("%s-%d", document.ID, len(chunks)),
			DocumentID: document.ID,
			Content:    content,
			Position:   len(chunks),
		}
		chunks = append(chunks, chunk)
	}
	return chunks, nil
}
