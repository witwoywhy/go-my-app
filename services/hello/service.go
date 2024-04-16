package hello

import (
	"myapp/libs/log"
)

type service struct {
}

func New() Service {
	return &service{}
}

func (s *service) Execute(request Request, l log.Logger) *Response {
	l.Info("This service is hello")
	ll := log.NewChildLog(l)
	ll.Info("This is span log")
	return &Response{
		Message: "world",
		AccNo: "AAAAAA",
	}
}
