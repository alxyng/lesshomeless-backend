package models

import "github.com/pborman/uuid"

type User struct {
	Id uuid.UUID `json:"id"`
}
