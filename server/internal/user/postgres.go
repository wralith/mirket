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
	query := `SELECT (id, username, email, hashed_password, about, created_at, updated_at) FROM "user"
	WHERE username=$1 AND deleted_at IS NULL`
	var user User

	err := r.db.QueryRow(ctx, query, username).Scan(&user)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (r *PostgresRepo) Get(ctx context.Context, id string) (*User, error) {
	query := `SELECT (id, username, email, hashed_password, about, created_at, updated_at) FROM "user"
	WHERE id=$1 AND deleted_at IS NULL`
	var user User

	err := r.db.QueryRow(ctx, query, id).Scan(&user)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (r *PostgresRepo) Delete(ctx context.Context, id string) error {
	query := `UPDATE "user" SET deleted_at = current_timestamp WHERE id=$1`

	if _, err := r.db.Exec(ctx, query, id); err != nil {
		return err
	}
	return nil
}

// Maybe different queries for each?

func (r *PostgresRepo) Update(ctx context.Context, user *User) error {
	query := `UPDATE "user"
	SET email = $1, hashed_password = $2, about = $3, updated_at = $4
	WHERE id = $5`

	if _, err := r.db.Exec(ctx, query, user.Email, user.HashedPassword, user.About, user.UpdatedAt, user.ID); err != nil {
		return err
	}
	return nil
}
