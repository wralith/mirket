package user

import (
	"context"
	"fmt"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/preset/postgres"
	"github.com/stretchr/testify/require"
)

var dummy = []*User{
	NewUser(NewUserOpts{Username: "test", Email: "test@mail.com", HashedPassword: []byte("secret-password")}),
}

func TestPostgresRepo(t *testing.T) {
	p := postgres.Preset(
		postgres.WithUser("root", "secret"),
		postgres.WithDatabase("mirket"),
		postgres.WithQueriesFile("../../migrations/user/000001_create_user_table.up.sql"),
	)

	container, err := gnomock.Start(p)
	require.NoError(t, err)

	defer func() { gnomock.Stop(container) }()

	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		container.Host, container.DefaultPort(), "root", "secret", "mirket",
	)

	pool, err := pgxpool.New(context.Background(), connStr)
	require.NoError(t, err)

	repo := NewPostgresRepo(pool)

	err = repo.Create(context.Background(), dummy[0])
	require.NoError(t, err)

	user, err := repo.GetByUsername(context.Background(), dummy[0].Username)
	require.NoError(t, err)
	require.Equal(t, user.Email, dummy[0].Email)
	require.Equal(t, user.Username, dummy[0].Username)

	err = repo.Delete(context.Background(), user.ID)
	require.NoError(t, err)

	_, err = repo.Get(context.Background(), user.ID)
	require.Error(t, err)
}
