package hello

import (
	"myapp/libs/log"
)

type Service interface {
	Execute(request Request, l log.Logger) *Response
}

type Request struct {
}

type Response struct {
	Message string `json:"message"`
	AccNo   string `json:"accNo"`
}
