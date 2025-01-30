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

	ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
