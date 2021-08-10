package domain

import "github.com/namrahov/banking/errs"

type Customer struct {
	Id          int    `db:"customer_id"`
	Name        string `db:"name"`
	City        string `db:"city"`
	Zipcode     string `db:"zipcode"`
	DateOfBirth string `db:"date_of_birth"`
	Status      string `db:"status"`
}

type CustomerRepository interface {
	FindAll() ([]Customer, *errs.AppError)
	FindById(string) (*Customer, *errs.AppError)
	FindAllByStatus(status string) ([]Customer, *errs.AppError)
}
