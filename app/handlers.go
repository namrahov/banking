package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/namrahov/banking/domain"
	"github.com/namrahov/banking/errhandler"
	"github.com/namrahov/banking/service"
	"io/ioutil"
	"log"
	"net/http"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) GetAllCustomers(w http.ResponseWriter, r *http.Request) {

	customers, err := ch.service.GetAllCustomer()

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(err.Code)
		json.NewEncoder(w).Encode(err.AsMessage())
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func (ch *CustomerHandlers) FindAllByStatus(w http.ResponseWriter, r *http.Request) {

	status := r.URL.Query().Get("status")

	customers, err := ch.service.FindAllByStatus(status)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(err.Code)
		json.NewEncoder(w).Encode(err.AsMessage())
	}

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

func (ch *CustomerHandlers) Save(w http.ResponseWriter, r *http.Request) {

	customer, err := getDtoFromRequest(r)
	if err != nil {
		return
	}

	customerDb, e := ch.service.Save(*customer)
	if e != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customerDb)
}

func (ch *CustomerHandlers) Update(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	ch.service.Update(name)
}

func getDtoFromRequest(r *http.Request) (*domain.Customer, error) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("failed to parse the model", b)
		return nil, errhandler.NewBadRequestError(err.Error())
	}

	customer := new(domain.Customer)
	err = json.Unmarshal(b, customer)
	if err != nil {
		log.Println("failed to parse the model", b)
		return nil, errhandler.NewBadRequestError(err.Error())
	}

	return customer, nil
}
