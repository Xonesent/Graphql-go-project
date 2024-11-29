// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

type GetUsersRequest struct {
	UserIds []UserId `json:"user_ids,omitempty"`
	Names   []string `json:"names,omitempty"`
}

type GetUsersResponse struct {
	Users []User `json:"users"`
}

type RegisterUserRequest struct {
	Name string `json:"name"`
}

type RegisterUserResponse struct {
	UserID UserId `json:"user_id"`
}