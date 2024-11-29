package user_usecase

import (
	"context"
	"graphql/internal/models"
	"graphql/internal/user/user_repository"
)

type UserPsqlRepo interface {
	CreateUser(ctx context.Context, userParams *user_repository.AddUser) (models.UserId, error)
	GetUsersByFilter(ctx context.Context, userFilter *user_repository.GetUserByFilter) ([]models.User, error)
}
