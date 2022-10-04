package usecase_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/wralith/mirket/src/userservice/internal/user/mock"
	"github.com/wralith/mirket/src/userservice/internal/user/repo"
	"github.com/wralith/mirket/src/userservice/internal/user/usecase"
	"github.com/wralith/mirket/src/userservice/internal/user/util"
	"github.com/wralith/mirket/src/userservice/pb"
)

func TestUseCase_AddUser(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockRepo(ctrl)

	uc := usecase.NewUserUsecase(mockRepo)

	user := &pb.User{
		Name:  util.RandomUsername(),
		Email: util.RandomEmail(),
	}

	mockRepo.EXPECT().AddUser(gomock.Any(), gomock.Eq(repo.AddUserParams{
		Name:  user.GetName(),
		Email: user.GetEmail(),
	})).Return(repo.User{
		Name:  user.GetName(),
		Email: user.GetEmail(),
	}, nil)

	newUser, err := uc.AddUser(user)
	require.NoError(t, err)
	require.NotNil(t, newUser)
	require.Equal(t, newUser.Id, user.Id)
	require.Equal(t, newUser.Name, user.Name)
	require.Equal(t, newUser.Email, user.Email)
	require.Equal(t, newUser.Bio, user.Bio)
}

func TestUseCase_GetUser(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockRepo(ctrl)

	uc := usecase.NewUserUsecase(mockRepo)

	user := &pb.User{
		Name:  util.RandomUsername(),
		Email: util.RandomEmail(),
	}

	mockRepo.EXPECT().AddUser(gomock.Any(), gomock.Any()).Return(repo.User{
		Name:  user.GetName(),
		Email: user.GetEmail(),
	}, nil)

	createdUser, err := uc.AddUser(user)
	require.NoError(t, err)

	mockRepo.EXPECT().GetUser(gomock.Any(), gomock.Eq(user.Id)).Return(repo.User{
		ID:    user.GetId(),
		Name:  user.GetName(),
		Email: user.GetEmail(),
	}, nil)

	queriedUser, err := uc.GetUser(createdUser.Id)
	require.NoError(t, err)

	require.Equal(t, queriedUser.Id, createdUser.Id)
	require.Equal(t, queriedUser.Name, createdUser.Name)
	require.Equal(t, queriedUser.Email, createdUser.Email)
	require.Equal(t, queriedUser.Bio, createdUser.Bio)
}
