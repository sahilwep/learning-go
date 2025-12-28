/*

// Why Repository Layer ?
	- Separates DB logic from HTTP logic.
	- Makes testing possible

*/

package user

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

/*
@Create
  - DB Query: storing email: passwordHash into DB
*/
func (r *Repository) Create(ctx context.Context, email, passwordHash string) error {
	query := `
		INSERT INTO users (email, password_hash)
		VALUES ($1, $2)
	`

	_, err := r.db.Exec(ctx, query, email, passwordHash)
	return err
}
