package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"bankapp/dto"
	"bankapp/service"
)

type AccountHandler struct {
	service service.AccountService
}

func (ah AccountHandler) createAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	customerId := vars["customer_id"]

	var request dto.NewAccountRequest

	// decode request into DTO
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerId = customerId
		
		account, appError := ah.service.NewAccount(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.Message)
		} else {
			writeResponse(w, http.StatusCreated, account)
		}
	}

}
