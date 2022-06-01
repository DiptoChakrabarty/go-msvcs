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
	return r.ErrCause
}

func NewRestError(msg string, status int, err string, cause []interface{}) RestErr {
	return restErr{
		ErrMessage: msg,
		ErrStatus:  status,
		ErrError:   err,
		ErrCause:   cause,
	}
}

func BadRequestError(m string) RestErr {
	return restErr{
		ErrMessage: m,
		ErrStatus:  http.StatusBadRequest,
		ErrError:   "bad_request",
	}
}

func NotFound(m string) RestErr {
	return restErr{
		ErrMessage: m,
		ErrStatus:  http.StatusNotFound,
		ErrError:   "not_found",
	}
}

func UnAuthorized(m string) RestErr {
	return restErr{
		ErrMessage: m,
		ErrStatus:  http.StatusUnauthorized,
		ErrError:   "unauthorized",
	}
}

func InternalServerError(m string, err error) RestErr {
	rerror := restErr{
		ErrMessage: m,
		ErrStatus:  http.StatusInternalServerError,
		ErrError:   "internal_error",
	}
	if err != nil {
		rerror.ErrCause = append(rerror.ErrCause, err)
	}
	return rerror
}
