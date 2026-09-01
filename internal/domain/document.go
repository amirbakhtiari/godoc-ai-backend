package domain

import "time"

type Document struct {
	ID        string
	Title     string
	Content   string
	Source    string
	CreatedAt time.Time
}
