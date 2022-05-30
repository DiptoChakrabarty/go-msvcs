package resterrors

import (
	"fmt"
	"net/http"
)

type RestErr interface {
	Message() string
	Status() int
	Error() string
	Causes() []interface{}
}

type restErr struct {
	Message string        `json:"msg"`
	Status  int           `json:"status"`
	ErrorErr   string        `json:error"`
	Cause   []interface{} `json:"causes"`
}

func (r restErr) Error() string {
	return fmt.Sprintf("message: %s - status: %d - error: %s - causes: %v",
		r.Message, r.Status, r.ErrorErr, r.Cause)
}






func BadRequestError(m string) *restErr {
	return &RestErr{
		Message: "Invalid Values Given",
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}
