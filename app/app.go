package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"bankapp/domain"
	"bankapp/service"
)

func Boot() {

	router := mux.NewRouter()

	// wiring: use stub
	// ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	// wiring: use db
	ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
