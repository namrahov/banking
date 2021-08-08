package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1", "Nurlan", "2000-01-01", "1"},
		{"2", "Ahmed", "2000-02-02", "1"},
	}

	return CustomerRepositoryStub{customers}
}
