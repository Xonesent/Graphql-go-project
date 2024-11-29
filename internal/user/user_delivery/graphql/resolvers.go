//go:generate go run github.com/99designs/gqlgen generate
package user_graphql

import (
	"context"
	"graphql/config"
	"graphql/internal/models"
	"graphql/internal/user/user_delivery"
)

type UserResolver struct {
	cfg    *config.Config
	userUC user_delivery.UserUC
}

func NewUserResolver(cfg *config.Config, userUC user_delivery.UserUC) *UserResolver {
	return &UserResolver{
		cfg:    cfg,
		userUC: userUC,
	}
}

func (r *mutationResolver) RegisterUser(ctx context.Context, request models.RegisterUserRequest) (*models.RegisterUserResponse, error) {
	userId, err := r.userUC.RegisterUser(ctx, toRegisterUser(request.Name))
	if err != nil {
		return nil, err
	}

	return toRegisterUserResponse(userId), nil
}

func (q *queryResolver) GetUsersByFilter(ctx context.Context, request models.GetUsersRequest) (*models.GetUsersResponse, error) {
	users, err := q.userUC.GetUserByFilter(ctx, toGetUserByFilter(request))
	if err != nil {
		return nil, err
	}

	return toGetUsersResponse(users), nil
}

func (r *UserResolver) Mutation() MutationResolver { return &mutationResolver{r} }

func (r *UserResolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *UserResolver }
type queryResolver struct{ *UserResolver }
