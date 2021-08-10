package dto

type CustomerResponse struct {
	Id     int    `json:"customerId"`
	Name   string `json:"fullName"`
	City   string `json:"city"`
	Status string `json:"status"`
}

type CustomerDto struct {
	Id          int    `json:"Id"`
	Name        string `json:"fullName"`
	City        string `json:"city"`
	Zipcode     string `json:"zipcode"`
	DateOfBirth string `json:"dateOfBirth"`
	Status      string `json:"status"`
}
