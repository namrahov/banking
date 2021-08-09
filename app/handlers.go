package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/namrahov/banking/service"
	"net/http"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) GetAllCustomers(w http.ResponseWriter, r *http.Request) {

	customers, _ := ch.service.GetAllCustomer()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func (ch *CustomerHandlers) FindById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idString := vars["customer_id"]

	customer, err := ch.service.FindById(idString)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(err.Code)
		json.NewEncoder(w).Encode(err.AsMessage())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customer)
}
