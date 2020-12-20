package app

import (
	"encoding/json"
	"net/http"

	"github.com/diogoqds/banking/dto"
	"github.com/diogoqds/banking/service"
	"github.com/gorilla/mux"
)

type AccountHandler struct {
	service service.AccountService
}

func (a AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerId = customerId
		account, appError := a.service.NewAccount(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.AsMessage())
		} else {
			writeResponse(w, http.StatusCreated, account)
		}
	}
}
