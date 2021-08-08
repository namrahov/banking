package main

import (
	"github.com/gorilla/mux"
	"github.com/namrahov/banking/app"
	"log"
	"net/http"
)

func main() {

	mux := mux.NewRouter()

	mux.HandleFunc("/greet", app.Greet).Methods(http.MethodGet)
	mux.HandleFunc("/customers", app.GetAllCustomers).Methods(http.MethodGet)
	mux.HandleFunc("/customer/{customer_id}", app.GetCustomerById).Methods(http.MethodGet)
	mux.HandleFunc("/customer", app.CreateCustomer).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
