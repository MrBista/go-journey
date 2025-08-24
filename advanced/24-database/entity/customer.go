package entity

type Customer struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func NewCustomer(id, name string) *Customer {

	return &Customer{
		Id:   id,
		Name: name,
	}
}
