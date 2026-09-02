package domain

import "time"

type Document struct {
	ID         string
	Title      string
	Content    string
	Source     string
	SourceType string

	CreatedAt time.Time
	UpdatedAt time.Time
}
