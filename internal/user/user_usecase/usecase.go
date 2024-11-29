package user_usecase

import (
	"context"
	"graphql/config"
	"graphql/internal/models"
)

type UserUseCase struct {
	cfg          *config.Config
	userPsqlRepo UserPsqlRepo
}

func NewUserUseCase(cfg *config.Config, userPsqlRepo UserPsqlRepo) *UserUseCase {
	return &UserUseCase{
		cfg:          cfg,
		userPsqlRepo: userPsqlRepo,
	}
}

func (u *UserUseCase) RegisterUser(ctx context.Context, registerUserParams *RegisterUser) (models.UserId, error) {
	userId, err := u.userPsqlRepo.CreateUser(ctx, registerUserParams.toCreateUser())
	if err != nil {
		return -1, err
	}

	return userId, nil
}

func (u *UserUseCase) GetUserByFilter(ctx context.Context, userFilter *GetUserByFilter) ([]models.User, error) {
	users, err := u.userPsqlRepo.GetUsersByFilter(ctx, userFilter.toGetUserByFilter())
	if err != nil {
		return nil, err
	}

	return users, nil
}
