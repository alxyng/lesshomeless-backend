package models

import "github.com/satori/uuid"

type User struct {
	Id uuid.UUID `json:"id"`
}
