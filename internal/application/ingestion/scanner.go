package ingestion

import (
	"os"
	"path/filepath"

	"github.com/amirbakhtiari/godoc-ai/internal/domain"
)

type Scanner struct {
	loader DocumentLoader
}

func NewScanner(loader DocumentLoader) *Scanner {
	return &Scanner{
		loader: loader,
	}
}
func (s *Scanner) Scan(root string) ([]domain.Document, error) {
	var documents []domain.Document

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) != ".md" {
			return nil
		}
		doc, err := s.loader.Load(path)
		if err != nil {
			return err
		}
		documents = append(documents, *doc)
		return nil
	})

	if err != nil {
		return nil, err
	}
	return documents, nil
}
