package model

type PersonRequest struct {
	PersonID  int    `json:"person_id"`
	LastName  string `json:"last_name" validate:"required"`
	FirstName string `json:"first_name" validate:"required"`
	Address   string `json:"address"`
	City      string `json:"city"`
}
