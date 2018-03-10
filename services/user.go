package services

import "github.com/nullseed/lesshomeless-backend/models"

type UserService interface {
	CreateUser() (models.User, error)
}
