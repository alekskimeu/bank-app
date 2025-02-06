package domain

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"bankapp/errs"
	"bankapp/logger"
)

type CustomerRepositoryDb struct {
	dbClient *sqlx.DB
}

func (db CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {

	var rows *sql.Rows
	var err error

	if status == "" {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		rows, err = db.dbClient.Query(findAllSql)
	} else {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		rows, err = db.dbClient.Query(findAllSql, status)
	}

	if err != nil {
		logger.LogError("Error fetching customers: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected DB error")
	}

	customers := make([]Customer, 0)

	err = sqlx.StructScan(rows, &customers)

	if err != nil {
		logger.LogError("Error scanning customers: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected DB error")
	}

	return customers, nil
}

func (db CustomerRepositoryDb) FindById(id string) (*Customer, *errs.AppError) {
	findCustomerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	row := db.dbClient.QueryRow(findCustomerSql, id)

	var customer Customer
	err := row.Scan(&customer.Id, &customer.Name, &customer.City, &customer.Zipcode, &customer.Dob, &customer.Status)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			logger.LogError("Error scanning customer: " + err.Error())
			return nil, errs.NewUnexpectedError("unexpected DB error")
		}
	}

	return &customer, nil
}

func NewCustomerRepositoryDb(dbClient *sqlx.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{dbClient}
}
