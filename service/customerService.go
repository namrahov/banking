package service

import (
	"github.com/namrahov/banking/domain"
	"github.com/namrahov/banking/dto"
	"github.com/namrahov/banking/errs"
	"github.com/namrahov/banking/mapper"
)

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, *errs.AppError)
	FindAllByStatus(status string) ([]domain.Customer, *errs.AppError)
	FindById(id string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, *errs.AppError) {
	return s.repo.FindAll()
}

func (s DefaultCustomerService) FindAllByStatus(status string) ([]domain.Customer, *errs.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	return s.repo.FindAllByStatus(status)
}

func (s DefaultCustomerService) FindById(id string) (*dto.CustomerResponse, *errs.AppError) {
	customer, err := s.repo.FindById(id)

	if err != nil {
		return nil, err
	}

	return mapper.MapCustomerToDto(customer), nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
