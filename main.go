package main

import (
	"github.com/namrahov/banking/tree/master/app"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/greet", app.Greet)
	http.HandleFunc("/customers", GetAllCustomers)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
