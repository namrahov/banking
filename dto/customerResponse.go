package dto

type CustomerResponse struct {
	Id     int    `json:"customerId"`
	Name   string `json:"fullName"`
	City   string `json:"city"`
	Status string `json:"status"`
}
