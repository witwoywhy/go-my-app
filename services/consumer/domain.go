package consumer

import "myapp/libs/log"

type Service interface {
	Execute(request Request, l log.Logger) (*Response, error)
}

type Request struct {
	Message string `json:"message"`
}

type Response struct {
}
