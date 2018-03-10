package user

import "github.com/nullseed/lesshomeless-backend/models"

type UserService interface {
	CreateUser() (*models.User, error)
	GetUser(userID string) (*models.User, error)
	AssignOfferToUser(*models.User, string) (*models.User, error)
}
