package user

import (
	"context"
	"fmt"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/preset/postgres"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
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

	defer func() {
		err := gnomock.Stop(container)
		require.NoError(t, err)
	}()

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
	require.Equal(t, dummy[0].Email, user.Email)
	require.Equal(t, dummy[0].Username, user.Username)

	user.UpdateAbout("test about")
	err = repo.Update(context.Background(), user)
	require.NoError(t, err)

	// reassign user
	user, err = repo.Get(context.Background(), user.ID)
	require.NoError(t, err)
	require.Equal(t, "test about", user.About)

	user.UpdateEmail("updated@mail.com")
	err = repo.Update(context.Background(), user)
	require.NoError(t, err)
	user, err = repo.Get(context.Background(), user.ID)
	require.NoError(t, err)
	require.Equal(t, "updated@mail.com", user.Email)

	hashed, err := bcrypt.GenerateFromPassword([]byte("updated"), bcrypt.DefaultCost)
	require.NoError(t, err)
	user.UpdatePassword(hashed)
	err = repo.Update(context.Background(), user)
	require.NoError(t, err)
	user, err = repo.Get(context.Background(), user.ID)
	require.NoError(t, err)
	err = bcrypt.CompareHashAndPassword(user.HashedPassword, []byte("updated"))
	require.NoError(t, err)

	err = repo.Delete(context.Background(), user.ID)
	require.NoError(t, err)

	_, err = repo.Get(context.Background(), user.ID)
	require.Error(t, err)
}
