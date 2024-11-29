package user_repository

import "graphql/internal/models"

type AddUser struct {
	Name string
}

type GetUserByFilter struct {
	UserIds []models.UserId
	Names   []string
}
