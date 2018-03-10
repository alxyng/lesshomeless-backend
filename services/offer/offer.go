package offer

import "github.com/nullseed/lesshomeless-backend/models"

type OfferService interface {
	// CreateOffer(models.Offer) (models.Offer, error)
	GetOffer(id string) (*models.Offer, error)
	// GetAllOffers() ([]models.Offer, error)
	// DeleteOffer(id string) error
}
