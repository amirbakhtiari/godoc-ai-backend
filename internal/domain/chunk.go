package domain

type Chunk struct {
	ID         string
	DocumentID string
	Content    string
	Position   int
	Metadata   ChunkMetadata
}
type ChunkMetadata struct {
	Title       string
	Section     string
	HeadingPath string
	Source      string
	SourceType  string
	ContentType string
}
