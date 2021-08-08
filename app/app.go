package app

import (
	"github.com/gorilla/mux"
	"github.com/namrahov/banking/domain"
	"github.com/namrahov/banking/service"
	"log"
	"net/http"
)

func Start() {
	router := mux.NewRouter()
	//ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	router.HandleFunc("/customers", ch.GetAllCustomers).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
