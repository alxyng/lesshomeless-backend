package user

import "github.com/nullseed/lesshomeless-backend/models"

type UserService interface {
	CreateUser() (*models.User, error)
	GetUser(userID string) (*models.User, error)
	AssignOfferToUser(*models.User, string) (*models.User, error)

	AssignReservationToUser(*models.User, string) (*models.User, error)
	RemoveReservationFromUser(*models.User, string) (*models.User, error)
	RemoveOfferFromUser(*models.User, string) (*models.User, error)
}
