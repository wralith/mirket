package user

import (
	"context"
	"time"

	"github.com/wralith/mirket/server/pkg/random"
)

type User struct {
	ID             string
	Username       string
	Email          string
	HashedPassword []byte
	About          string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type NewUserOpts struct {
	Username       string
	Email          string
	HashedPassword []byte
}

func NewUser(opts NewUserOpts) *User {
	now := time.Now()
	return &User{
		ID:             random.ID(),
		Username:       opts.Username,
		Email:          opts.Email,
		HashedPassword: opts.HashedPassword,
		About:          "",
		CreatedAt:      now,
		UpdatedAt:      now,
		DeletedAt:      nil,
	}
}

type Repo interface {
	Create(context.Context, *User) error
	GetByUsername(ctx context.Context, username string) (*User, error)
}
