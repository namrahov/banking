package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/greet", Greet)
	http.HandleFunc("/customers", GetAllCustomers)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
