package models

type Me struct {
	Giving   Offer `json:"giving"`
	Reserved Offer `json:"reserved"`
}
