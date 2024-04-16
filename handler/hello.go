package handler

import (
	"myapp/libs/app"
	"myapp/libs/log"
	"myapp/services/hello"
)

type helloHandler struct {
	svc hello.Service
}

func NewHelloHandler(svc hello.Service) *helloHandler {
	return &helloHandler{
		svc: svc,
	}
}

func (h *helloHandler) Handle(ctx app.WebFrameworkContext, l log.Logger) error {
	resp := h.svc.Execute(hello.Request{}, l)
	return ctx.JSON(200, resp)
}
