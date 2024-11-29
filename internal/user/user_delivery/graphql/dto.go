package user_graphql

import (
	"graphql/internal/models"
	"graphql/internal/user/user_usecase"
)

func toRegisterUser(name string) *user_usecase.RegisterUser {
	return &user_usecase.RegisterUser{
		Name: name,
	}
}

func toRegisterUserResponse(userId models.UserId) *models.RegisterUserResponse {
	return &models.RegisterUserResponse{
		UserID: userId,
	}
}

func toGetUserByFilter(request models.GetUsersRequest) *user_usecase.GetUserByFilter {
	return &user_usecase.GetUserByFilter{
		UserIds: request.UserIds,
		Names:   request.Names,
	}
}

func toGetUsersResponse(users []models.User) *models.GetUsersResponse {
	return &models.GetUsersResponse{
		Users: users,
	}
}
