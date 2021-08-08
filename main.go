package main

import (
	"github.com/namrahov/banking/app"
	"log"
	"net/http"
)

func main() {

	//mux := NewHandlerMux()

	http.HandleFunc("/greet", app.Greet)
	http.HandleFunc("/customers", app.GetAllCustomers)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
