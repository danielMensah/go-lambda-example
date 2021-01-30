package main

import (
	"github.com/danielMensah/go-lambda-example/internal/repository"
	"github.com/danielMensah/go-lambda-example/internal/request"
	"github.com/danielMensah/go-lambda-example/internal/utility/aws"
)

func GetCustomers(req aws.Request) aws.Response {
	var data request.GetCustomersRequest

	if err := aws.Decode([]byte(req.Body), &data); err != nil {
		return aws.ErrorResponse(aws.InternalServerError.SetError(err), "decoding request")
	}

	repo := repository.NewRepository()

	customers, getCustomersError := repo.CustomerRepo.GetCustomers(data.Store)

	if getCustomersError != nil {
		return aws.ErrorResponse(aws.BadRequestError.SetError(getCustomersError), getCustomersError.Error())
	}

	return aws.JsonResponse(customers)
}
