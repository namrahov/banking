package service

type CustomerService interface {
	GetAllCustomer() ([]Customer, error)
}
