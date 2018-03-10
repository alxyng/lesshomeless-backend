package models

import (
	"time"
)

type Offer struct {
	Id          string       `json:"id"`
	Name        string       `json:"name"`
	CreatedOn   time.Time    `json:"createdOn"`
	CreatedBy   string       `json:"createdBy"`
	Location    Location     `json:"location"`
	Reservation *Reservation `json:"reservation,omitempty"`
}

type Location struct {
	Lat  float32 `json:"lat"`
	Long float32 `json:"long"`
}

type Reservation struct {
	ReservedBy   string    `json:"reservedBy"`
	ReservedOn   time.Time `json:"reservedOn"`
	Acknowledged bool      `json:"acknowledged"`
}
