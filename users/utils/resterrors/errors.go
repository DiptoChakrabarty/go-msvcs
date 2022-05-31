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
	ErrMessage string        `json:"msg"`
	ErrStatus  int           `json:"status"`
	ErrError   string        `json:error"`
	ErrCause   []interface{} `json:"causes"`
}

func (r restErr) Error() string {
	return fmt.Sprintf("message: %s - status: %d - error: %s - causes: %v",
		r.ErrMessage, r.ErrStatus, r.ErrError, r.ErrCause)
}

func (r restErr) Message() string {
	return r.ErrMessage
}

func (r restErr) Status() int {
	return r.ErrStatus
}

func (r restErr) Causes() []interface{} {
	return r.Cause
}

func BadRequestError(m string) *restErr {
	return &restErr{
		ErrMessage: "Invalid Values Given",
		ErrStatus:  http.StatusBadRequest,
		ErrError:   "bad_request",
	}
}
