package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
}

func (db CustomerRepositoryDb) FindAll() ([]Customer, error) {

	dbClient, err := sql.Open("mysql", "root:Soda3291@/banking")
	if err != nil {
		panic(err)
	}

	// See "Important settings" section.
	dbClient.SetConnMaxLifetime(time.Minute * 3)
	dbClient.SetMaxOpenConns(10)
	dbClient.SetMaxIdleConns(10)

	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"

	rows, err := dbClient.Query(findAllSql)

	if err != nil {
		log.Println("Error fetching customers: ", err.Error())
	}

	customers := make([]Customer, 0)

	for rows.Next() {
		var customer Customer

		err := rows.Scan(&customer.Id, &customer.Name, &customer.City, &customer.Zipcode, &customer.Dob, &customer.Status)

		if err != nil {
			log.Println("Error scanning customers: ", err.Error())
		}

		customers = append(customers, customer)
	}

	return customers, nil
}
