package domain

type Customer struct {
	Id          string
	Name        string
	DateofBirth string
	Status      string
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
}
