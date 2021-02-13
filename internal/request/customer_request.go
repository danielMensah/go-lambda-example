package request

import (
	"encoding/json"
	validator "github.com/asaskevich/govalidator"
)

type GetCustomersRequest struct {
	Store string `json:"store" valid:"required~store required"`
}

func (r GetCustomersRequest) Decode(req []byte) error {
	err := json.Unmarshal(req, &r)

	if err != nil {
		return err
	}

	return r.valid()
}

func (r GetCustomersRequest) valid() error {
	if _, err := validator.ValidateStruct(r); err != nil {
		return err
	}

	return nil
}