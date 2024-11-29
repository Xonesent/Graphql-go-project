package user_delivery

import (
	"context"
	"graphql/internal/models"
	"graphql/internal/user/user_usecase"
)

type UserUC interface {
	RegisterUser(ctx context.Context, registerUserParams *user_usecase.RegisterUser) (models.UserId, error)
	GetUserByFilter(ctx context.Context, userFilter *user_usecase.GetUserByFilter) ([]models.User, error)
}
