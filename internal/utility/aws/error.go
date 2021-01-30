package aws

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Error struct {
	StatusCode int
	Error error
}

var (
	BadRequestError      = NewError(http.StatusBadRequest, nil)
	ConflictError        = NewError(http.StatusConflict, nil)
	InternalServerError  = NewError(http.StatusInternalServerError, nil)
	NotFoundError        = NewError(http.StatusNotFound, nil)
	TooManyRequestsError = NewError(http.StatusTooManyRequests, nil)
	UnauthorizedError    = NewError(http.StatusUnauthorized, nil)
)

func NewError(status int, err error) *Error {

	e := &Error{}

	if err == nil {
		err = errors.New(http.StatusText(status))
	}

	e.SetStatus(status).SetError(err)

	return e
}

func (e *Error) SetError(err error) *Error {
	e.Error = err
	return e
}

func (e *Error) SetStatus(status int) *Error {
	e.StatusCode = status
	return e
}

func (e *Error) ToJson() string {
	var errArr []string
	var response = map[string]*[]string{
		"error": &errArr,
	}

	if err := json.Unmarshal([]byte(e.Error.Error()), &errArr); err != nil {
		errArr = append(errArr, e.Error.Error())
	}

	jsonResponse, _ := json.Marshal(response)

	return string(jsonResponse)
}