package models

type Users struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string`json:"email"`
	CreatedAt string `json:"createdat"`
	Active bool `json:"active"`
}