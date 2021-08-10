package domain

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/namrahov/banking/errs"
	"log"
)

type CustomerRepositoryDb struct {
	conn *sql.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, *errs.AppError) {

	rows, err := d.conn.Query("select city, customer_id, date_of_birth, name, zipcode, status from customers")

	if err != nil {
		log.Println("Error while querying customer table " + err.Error())
		return nil, errs.NewNotFoundError("Customers not found")
	}

	defer rows.Close()

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.City, &c.Id, &c.DateOfBirth, &c.Name, &c.Zipcode, &c.Status)

		if err != nil {
			if err == sql.ErrNoRows {
				return nil, errs.NewNotFoundError("Customer not found")
			} else {
				log.Println("Error while scanning customers " + err.Error())
				return nil, errs.NewUnexpectedError("Unexpected database error")
			}
		}

		customers = append(customers, c)
	}

	return customers, nil
}

func (d CustomerRepositoryDb) FindAllByStatus(status string) ([]Customer, *errs.AppError) {

	rows, err := d.conn.Query("select city, customer_id, date_of_birth, name, zipcode, status from customers where status = $1", status)

	if err != nil {
		log.Println("Error while querying customer table " + err.Error())
		return nil, errs.NewNotFoundError("Customers not found")
	}

	defer rows.Close()

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.City, &c.Id, &c.DateOfBirth, &c.Name, &c.Zipcode, &c.Status)

		if err != nil {
			if err == sql.ErrNoRows {
				return nil, errs.NewNotFoundError("Customer not found")
			} else {
				log.Println("Error while scanning customers " + err.Error())
				return nil, errs.NewUnexpectedError("Unexpected database error")
			}
		}

		customers = append(customers, c)
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

func (d CustomerRepositoryDb) Save(c Customer) (*Customer, *errs.AppError) {
	sqlInsert := "INSERT INTO CUSTOMERS (city, date_of_birth, name, zipcode, status) values ($1, $2 , $3, $4, $5) RETURNING customer_id"

	result, err := d.conn.Exec(sqlInsert, c.City, c.DateOfBirth, c.Name, c.Zipcode, c.Status)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			log.Println("Error while save customer " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	id, err := result.LastInsertId()
	c.Id = int(id)

	return &c, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	conn, err := sql.Open("pgx", "host=localhost port=5432 dbname=postgres user=postgres password=boot")

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
