package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"bankapp/service"
)

type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getAllCustomers(w http.ResponseWriter, r *http.Request) {

	customers, err := ch.service.GetAllCustomers()

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	}

	writeResponse(w, http.StatusOK, customers)
}

func (ch *CustomerHandler) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["id"]

	customer, err := ch.service.GetCustomer(customerId)

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	}
	writeResponse(w, http.StatusOK, customer)
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
