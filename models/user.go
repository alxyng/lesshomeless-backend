package models

type User struct {
	Id       string   `json:"id"`
	Offers   []string `json:"offers"`
	Reserved []string `json:"reserved"`
}
