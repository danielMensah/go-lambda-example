package main

import (
	"github.com/danielMensah/go-lambda-example/internal/repository"
	"github.com/danielMensah/go-lambda-example/internal/request"
	"github.com/danielMensah/go-lambda-example/internal/utility/aws"
)

func GetOrders(req aws.Request) aws.Response {
	var data request.GetOrdersRequest

	if err := aws.Decode([]byte(req.Body), &data); err != nil {
		return aws.ErrorResponse(aws.BadRequestError.SetError(err), "decoding request data")
	}

	repo := repository.NewRepository()

	orders, getOrdersErr := repo.OrderRepo.GetOrders(data.CustomerId)

	if getOrdersErr != nil {
		return aws.ErrorResponse(aws.BadRequestError.SetError(getOrdersErr), "getting orders")
	}

	return aws.JsonResponse(orders)
}
