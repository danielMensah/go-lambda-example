package request

import (
	"encoding/json"
	validator "github.com/asaskevich/govalidator"
)

type GetOrdersRequest struct {
	CustomerId string `json:"customerId" valid:"required~customer id required"`
}

func (r GetOrdersRequest) Decode(req []byte) error {
	err := json.Unmarshal(req, &r)

	if err != nil {
		return err
	}

	return r.valid()
}

func (r GetOrdersRequest) valid() error {
	if _, err := validator.ValidateStruct(r); err != nil {
		return err
	}

	return nil
}