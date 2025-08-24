package entity

import "time"

type Customer struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Balance   int32     `json:"balance"`
	Rating    float64   `json:"rating"`
	CreatedAt time.Time `json:"createdAt"`
	BirthDate time.Time `json:"birthDate"`
	Married   bool      `json:"married"`
}

func NewCustomer(id, name string) *Customer {

	return &Customer{
		Id:   id,
		Name: name,
	}
}
