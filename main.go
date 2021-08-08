package main

import (
	"github.com/gorilla/mux"
	"github.com/namrahov/banking/app"
	"log"
	"net/http"
)

func main() {

	mux := mux.NewRouter()

	mux.HandleFunc("/greet", app.Greet)
	mux.HandleFunc("/customers", app.GetAllCustomers)
	mux.HandleFunc("/customer/{customer_id}", app.GetCustomerById)

	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
