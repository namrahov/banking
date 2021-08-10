package mapper

import (
	"github.com/namrahov/banking/domain"
	"github.com/namrahov/banking/dto"
)

func MapCustomerToDto(customer *domain.Customer) *dto.CustomerResponse {

	return &dto.CustomerResponse{
		Id:     customer.Id,
		Name:   customer.Name,
		City:   customer.City,
		Status: customer.Status,
	}
}
