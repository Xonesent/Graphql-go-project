package user_http

import (
	"graphql/internal/models"
	"graphql/internal/user/user_usecase"
)

type registerUserRequest struct {
	Name string `json:"name" validate:"required"`
}

func (r *registerUserRequest) toRegisterUser() *user_usecase.RegisterUser {
	return &user_usecase.RegisterUser{
		Name: r.Name,
	}
}

type getUserRequest struct {
	UserIds []models.UserId `json:"ids"`
	Names   []string        `json:"names"`
}

func (r *getUserRequest) toGetUserByFilter() *user_usecase.GetUserByFilter {
	return &user_usecase.GetUserByFilter{
		UserIds: r.UserIds,
		Names:   r.Names,
	}
}
