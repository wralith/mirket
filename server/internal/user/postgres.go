package user

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepo struct {
	db *pgxpool.Pool
}

func NewPostgresRepo(db *pgxpool.Pool) *PostgresRepo {
	return &PostgresRepo{
		db: db,
	}
}

var _ Repo = &PostgresRepo{}

func (r *PostgresRepo) Create(ctx context.Context, user *User) error {
	query := `INSERT INTO "user" (id, username, email, hashed_password, about) VALUES ($1, $2, $3, $4, $5)`
	if _, err := r.db.Exec(ctx, query, user.ID, user.Username, user.Email, user.HashedPassword, user.About); err != nil {
		return err
	}
	return nil
}

func (r *PostgresRepo) GetByUsername(ctx context.Context, username string) (*User, error) {
	query := `SELECT (id, username, email, hashed_password, about, created_at, updated_at) FROM "user" WHERE username=$1`
	var user User

	err := r.db.QueryRow(ctx, query, username).Scan(&user)
	if err != nil {
		return nil, err
	}
	return &user, err
}
