package models

import (
	"time"

	"github.com/satori/uuid"
)

type Offer struct {
	Id          uuid.UUID   `json:"id"`
	Name        string      `json:"name"`
	CreatedOn   time.Time   `json:"createdOn"`
	CreatedBy   uuid.UUID   `json:"createdBy"`
	Location    Location    `json:"location"`
	Reservation Reservation `json:"reservation"`
}

type Location struct {
	Lat  float32 `json:"lat"`
	Long float32 `json:"long"`
}

type Reservation struct {
	ReservedBy   uuid.UUID `json:"reservedBy"`
	ReservedOn   time.Time `json:"reservedOn"`
	Acknowledged bool      `json:"acknowledged"`
}
