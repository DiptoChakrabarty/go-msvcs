package resterrors

import (
	"net/http"
)

type RestErr struct {
	Message string `json:"msg"`
	Status  int    `json:"status"`
	Error   string `json:error"`
}

func BadRequestError(m string) *RestErr {
	return &RestErr{
		Message: "Invalid Values Given",
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}
