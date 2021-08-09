package app

import (
	"encoding/json"
	"fmt"
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
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}
