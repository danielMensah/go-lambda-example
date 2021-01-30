package repository

import (
	"errors"
	faker "github.com/brianvoe/gofakeit"
)

var GetOrdersError = errors.New("no customer id provided")

type Order struct {
	Id string
	Name string
	Email string
	Phone string
}

func (o Order) GetOrders(customerId string) ([]Order, error) {
	if customerId == "" {
		return nil, GetOrdersError
	}

	orders := generateOrders(5)

	return orders, nil
}

func generateOrders(n int) []Order {
	var orders []Order

	for i := 0; i < n; i++ {
		orders = append(orders, Order{
			Id: faker.UUID(),
			Name: faker.Name(),
			Email: faker.Email(),
			Phone: faker.Phone(),
		})
	}

	return orders
}