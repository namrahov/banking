package domain

import "github.com/namrahov/banking/errs"

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, *errs.AppError) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{1, "Nurlan", "Sumgait", "AZ5011", "2000-01-01", "1"},
		{2, "Ahmed", "Baki", "AZ5000", "2000-02-02", "1"},
	}

	return CustomerRepositoryStub{customers}
}
