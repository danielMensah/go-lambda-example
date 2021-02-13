package repository

import (
	faker "github.com/brianvoe/gofakeit"
)

type Order struct {
	Id string
	Email string
	Amount float64
	Date string
}

func (o Order) GetOrders(customerId string) ([]Order, error) {
	orders := generateOrders(5)

	return orders, nil
}

func generateOrders(n int) []Order {
	var orders []Order

	for i := 0; i < n; i++ {
		orders = append(orders, Order{
			Id: faker.UUID(),
			Email: faker.Email(),
			Amount: faker.Price(1000, 10000),
			Date: faker.Date().String(),
		})
	}

	return orders
}