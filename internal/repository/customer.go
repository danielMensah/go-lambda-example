package repository

import (
	"errors"
	faker "github.com/brianvoe/gofakeit"
)

var GetCustomerError = errors.New("no store provided")

type Customer struct {
	Id string
	Name string
	Email string
	Phone string
}

func (c Customer) GetCustomers(store string) ([]Customer, error) {
	if store == "" {
		return nil, GetCustomerError
	}

	customers := generateCustomers(5)

	return customers, nil
}

func generateCustomers(n int) []Customer {
	var customers []Customer

	for i := 0; i < n; i++ {
		customers = append(customers, Customer{
			Id: faker.UUID(),
			Name: faker.Name(),
			Email: faker.Email(),
			Phone: faker.Phone(),
		})
	}

	return customers
}