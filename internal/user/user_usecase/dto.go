package user_usecase

import (
	"graphql/internal/models"
	"graphql/internal/user/user_repository"
)

type RegisterUser struct {
	Name string
}

func (r *RegisterUser) toCreateUser() *user_repository.AddUser {
	return &user_repository.AddUser{
		Name: r.Name,
	}
}

type GetUserByFilter struct {
	UserIds []models.UserId
	Names   []string
}

func (r *GetUserByFilter) toGetUserByFilter() *user_repository.GetUserByFilter {
	return &user_repository.GetUserByFilter{
		UserIds: r.UserIds,
		Names:   r.Names,
	}
}
