package models

type User struct {
	UserId UserId `json:"user_id"`
	Name   string `json:"name"`
}
