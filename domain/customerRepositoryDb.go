package domain

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/namrahov/banking/errs"
	"log"
)

type CustomerRepositoryDb struct {
	conn *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, *errs.AppError) {

	var err error
	customers := make([]Customer, 0)

	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
	err = d.conn.Select(&customers, findAllSql)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			log.Println("Error while scanning customers " + err.Error()) //log.Error("Error while scanning customers " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return customers, nil
}

func (d CustomerRepositoryDb) FindAllByStatus(status string) ([]Customer, *errs.AppError) {

	var err error
	customers := make([]Customer, 0)

	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = $1"
	err = d.conn.Select(&customers, findAllSql, status)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			log.Println("Error while scanning customers " + err.Error()) //log.Error("Error while scanning customers " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return customers, nil
}

func (d CustomerRepositoryDb) FindById(id string) (*Customer, *errs.AppError) {
	customerSql := "select city, customer_id, date_of_birth, name, zipcode, status from customers where customer_id = $1"
	row := d.conn.QueryRow(customerSql, id)

	var c Customer
	err := row.Scan(&c.City, &c.Id, &c.DateOfBirth, &c.Name, &c.Zipcode, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			log.Println("Error while scanning customers " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &c, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	conn, err := sqlx.Open("pgx", "host=localhost port=5432 dbname=postgres user=postgres password=boot")

	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to connect:%v\n", err))
	}

	log.Println("Connected to database!")

	// test my connection
	err = conn.Ping()
	if err != nil {
		log.Fatal("Cannot ping database!")
	}

	log.Println("Pinged database!")

	return CustomerRepositoryDb{conn}
}
