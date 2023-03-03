package user

import (
	"context"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo Repo
}

func NewService(repo Repo) *Service {
	return &Service{repo: repo}
}

// UserResponse represent user information that will be sent to outside of this module
type UserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	About    string `json:"about"`

	CreatedAt time.Time `json:"createdAt"`
}

func ToUserResponse(u *User) UserResponse {
	return UserResponse{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		About:     u.About,
		CreatedAt: u.CreatedAt,
	}
}

func (s *Service) Create(ctx context.Context, dto NewUserOpts) error {
	user := NewUser(dto)
	return s.repo.Create(ctx, user)
}

func (s *Service) GetByUsername(ctx context.Context, username string) (UserResponse, error) {
	user, err := s.repo.GetByUsername(ctx, username)
	return ToUserResponse(user), err
}

func (s *Service) Get(ctx context.Context, id string) (UserResponse, error) {
	user, err := s.repo.Get(ctx, id)
	return ToUserResponse(user), err
}

func (s *Service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *Service) CheckPassword(ctx context.Context, username string, password string) error {
	user, err := s.repo.GetByUsername(ctx, username)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password))
	if err != nil {
		return errors.New("invalid credentials")
	}
	return nil
}
