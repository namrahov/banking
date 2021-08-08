package main

import (
	"github.com/gorilla/mux"
	"github.com/namrahov/banking/app"
	"github.com/namrahov/banking/domain"
	"github.com/namrahov/banking/service"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter()

	ch := app.CustomerHandlers{
		service.NewCustomerService(domain.NewCustomerRepositoryStub()),
	}
	router.HandleFunc("/customers", ch.GetAllCustomers).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
