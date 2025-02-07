package app

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"bankapp/domain"
	"bankapp/service"
)

func Boot() {

	router := mux.NewRouter()

	// wiring: use stub
	// ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	dbClient := getDbClient()

	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)

	// wiring: use db
	ch := CustomerHandler{service.NewCustomerService(customerRepositoryDb)}
	ah := AccountHandler{service.NewAccountService(accountRepositoryDb)}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.createAccount).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", ah.makeTransfer).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe("localhost:8080", router))
}

func getDbClient() *sqlx.DB {
	dbClient, err := sqlx.Open("mysql", "root:Soda3291@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}

	dbClient.SetConnMaxLifetime(time.Minute * 3)
	dbClient.SetMaxOpenConns(10)
	dbClient.SetMaxIdleConns(10)

	return dbClient
}
