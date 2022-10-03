package usecase

import (
	"context"

	"github.com/wralith/mirket/src/userservice/internal/user/repo"
	"github.com/wralith/mirket/src/userservice/pb"
)

type UserUsecase struct {
	repo repo.Repo
}

func NewUserUsecase(repo repo.Repo) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (u *UserUsecase) AddUser(pr *pb.User) (*pb.User, error) {
	arg := repo.AddUserParams(repo.AddUserParams{Name: pr.Name, Email: pr.Email})

	res, err := u.repo.AddUser(context.Background(), arg)
	if err != nil {
		return nil, err
	}

	pbUser := pb.User{Id: res.ID, Name: res.Name, Email: res.Email, Bio: res.Bio}
	return &pbUser, nil
}

func (u *UserUsecase) GetUser(id uint32) (*pb.User, error) {
	res, err := u.repo.GetUser(context.Background(), id)
	if err != nil {
		return nil, err
	}

	pbUser := pb.User{Id: res.ID, Name: res.Name, Email: res.Email, Bio: res.Bio}
	return &pbUser, nil
}
