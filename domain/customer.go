package domain

type Customer struct {
	Id          int    `json:"customer_id"`
	Name        string `json:"full_name"`
	City        string `json:"city"`
	Zipcode     string `json:"zipcode"`
	DateOfBirth string `json:"date_of_birth"`
	Status      string `json:"status"`
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	FindById(string) (*Customer, error)
}
