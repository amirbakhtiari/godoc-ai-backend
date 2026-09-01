package documents

import (
	"context"

	"github.com/amirbakhtiari/godoc-ai/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepository struct {
	db *pgxpool.Pool
}

func NewPostgresRepository(db *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}
func (r *PostgresRepository) Create(ctx context.Context, document *domain.Document) error {
	_, err := r.db.Exec(ctx,
		`INSERT INTO documents (
		id,
		title,
		content,
		source,
		created_at
	)
	VALUES ($1, $2, $3, $4, $5)`,
		document.ID,
		document.Title,
		document.Content,
		document.Source,
		document.CreatedAt,
	)
	return err
}
