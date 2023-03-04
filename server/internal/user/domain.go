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
	Get(ctx context.Context, id string) (*User, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, user *User) error
}

func (u *User) UpdateEmail(email string) {
	u.Email = email
	u.UpdatedAt = time.Now()
}

func (u *User) UpdatePassword(hashedPassword []byte) {
	u.HashedPassword = hashedPassword
	u.UpdatedAt = time.Now()
}

func (u *User) UpdateAbout(about string) {
	u.About = about
	u.UpdatedAt = time.Now()
}
