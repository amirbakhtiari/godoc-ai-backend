package document

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/amirbakhtiari/godoc-ai/internal/domain"
)

type MarkdownLoader struct{}

func NewMarkdownLoader() *MarkdownLoader {
	return &MarkdownLoader{}
}
func (loader *MarkdownLoader) Load(path string) (document *domain.Document, err error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	content := string(data)

	title := extractTitle(content)

	return &domain.Document{
		ID:        filepath.Base(path),
		Title:     title,
		Content:   content,
		Source:    path,
		CreatedAt: time.Time{},
	}, nil
}

func extractTitle(content string) string {
	for _, line := range strings.Split(content, "\n") {
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "# ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "# "))
		}
	}

	return "Untitled"
}
