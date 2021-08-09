package domain

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
)

type CustomerRepositoryDb struct {
	conn *sql.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {

	rows, err := d.conn.Query("select city, customer_id, date_of_birth, name, zipcode, status from customers")

	if err != nil {
		log.Println("Error while querying customer table " + err.Error())
		return nil, err
	}

	defer rows.Close()

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.City, &c.Id, &c.DateOfBirth, &c.Name, &c.Zipcode, &c.Status)
		if err != nil {
			log.Println("Error while scanning customers " + err.Error())
			return nil, err
		}
		customers = append(customers, c)
	}

	return customers, nil
}

func (d CustomerRepositoryDb) FindById(id string) (*Customer, error) {
	customerSql := "select city, customer_id, date_of_birth, name, zipcode, status from customers where customer_id = $1"
	row := d.conn.QueryRow(customerSql, id)

	var c Customer
	err := row.Scan(&c.City, &c.Id, &c.DateOfBirth, &c.Name, &c.Zipcode, &c.Status)
	if err != nil {
		log.Println("Error while scanning customers " + err.Error())
		return nil, err
	}

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
