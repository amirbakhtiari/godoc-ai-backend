package documents

import (
	"fmt"
	"strings"

	"github.com/amirbakhtiari/godoc-ai/internal/domain"
)

type MarkdownChunker struct {
	maxWords int
}

func NewMarkdownChunker(maxWords int) *MarkdownChunker {
	if maxWords <= 0 {
		maxWords = 500
	}

	return &MarkdownChunker{
		maxWords: maxWords,
	}
}

func (c MarkdownChunker) Chunk(
	document *domain.Document,
) ([]domain.Chunk, error) {

	sections := parseMarkdownSections(document.Content)

	var chunks []domain.Chunk

	for _, section := range sections {
		content := strings.TrimSpace(section.Content)

		if content == "" {
			continue
		}

		sectionChunks := c.chunkSection(
			document,
			section,
		)

		chunks = append(chunks, sectionChunks...)
	}

	for i := range chunks {
		chunks[i].Position = i
		chunks[i].ID = fmt.Sprintf(
			"%s-%d",
			document.ID,
			i,
		)
	}

	return chunks, nil
}

func (c MarkdownChunker) chunkSection(
	document *domain.Document,
	section markdownSection,
) []domain.Chunk {

	if wordCount(section.Content) <= c.maxWords {
		return []domain.Chunk{
			c.newChunk(
				document,
				section,
				section.Content,
			),
		}
	}

	blocks := splitMarkdownBlocks(section.Content)

	var chunks []domain.Chunk
	var currentBlocks []string
	currentWords := 0

	flush := func() {
		if len(currentBlocks) == 0 {
			return
		}

		content := strings.TrimSpace(
			strings.Join(currentBlocks, "\n\n"),
		)

		if content == "" {
			currentBlocks = nil
			currentWords = 0
			return
		}

		chunks = append(
			chunks,
			c.newChunk(
				document,
				section,
				content,
			),
		)

		currentBlocks = nil
		currentWords = 0
	}

	for _, block := range blocks {
		block = strings.TrimSpace(block)

		if block == "" {
			continue
		}

		blockWords := wordCount(block)

		// A code block or any single block larger than maxWords
		// stays intact. We don't want to destroy semantic structure.
		if blockWords > c.maxWords {
			flush()

			chunks = append(
				chunks,
				c.newChunk(
					document,
					section,
					block,
				),
			)

			continue
		}

		if currentWords+blockWords > c.maxWords {
			flush()
		}

		currentBlocks = append(currentBlocks, block)
		currentWords += blockWords
	}

	flush()

	return chunks
}

func (c MarkdownChunker) newChunk(
	document *domain.Document,
	section markdownSection,
	content string,
) domain.Chunk {

	return domain.Chunk{
		DocumentID: document.ID,
		Content:    strings.TrimSpace(content),
		Metadata: domain.ChunkMetadata{
			Title:       document.Title,
			Section:     section.Title,
			HeadingPath: section.HeadingPath,
			Source:      document.Source,
			SourceType:  document.SourceType,
			ContentType: detectContentType(content),
		},
	}
}

func detectContentType(content string) string {
	trimmed := strings.TrimSpace(content)

	if strings.HasPrefix(trimmed, "```") ||
		strings.HasPrefix(trimmed, "~~~") {
		return "code"
	}

	return "text"
}

func wordCount(content string) int {
	return len(strings.Fields(content))
}

func splitMarkdownBlocks(content string) []string {
	lines := strings.Split(content, "\n")

	var blocks []string
	var current []string

	inCodeBlock := false

	flush := func() {
		block := strings.TrimSpace(
			strings.Join(current, "\n"),
		)

		if block != "" {
			blocks = append(blocks, block)
		}

		current = nil
	}

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		if isCodeFence(trimmed) {
			current = append(current, line)
			inCodeBlock = !inCodeBlock
			continue
		}

		if inCodeBlock {
			current = append(current, line)
			continue
		}

		// Empty line marks a block boundary.
		if trimmed == "" {
			flush()
			continue
		}

		current = append(current, line)
	}

	flush()

	return blocks
}
