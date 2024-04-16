package handler

import (
	"myapp/libs/app"
	"myapp/libs/log"
	"myapp/services/publish"
)

type publishHandler struct {
	svc publish.Service
}

func NewPublishHandler(svc publish.Service) *publishHandler {
	return &publishHandler{
		svc: svc,
	}
}

func (h *publishHandler) Handle(ctx app.WebFrameworkContext, l log.Logger) error {
	var req publish.Request

	if err := ctx.BindJSON(&req); err != nil {
		l.Errorf("Failed to bind json err: %v", err)
		return ctx.JSON(400, nil)
	}

	resp, err := h.svc.Execute(req, l)
	if err != nil {
		return ctx.JSON(500, nil)
	}

	return ctx.JSON(200, resp)
}
