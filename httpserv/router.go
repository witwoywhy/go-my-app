package httpserv

import (
	"myapp/handler"
	"myapp/infra"
	"myapp/libs/app"
	m "myapp/libs/app"
	"myapp/services/hello"
	"myapp/services/publish"
	"net/http"
)

func bindHelloHandler(app app.App) {
	h := handler.NewHelloHandler(hello.New())
	app.Register(
		http.MethodGet,
		"/",
		m.UseHandle(h.Handle),
	)
}

func bindPublishHandler(app app.App) {
	svc := publish.New(infra.Writer)
	h := handler.NewPublishHandler(svc)
	app.Register(
		http.MethodPost,
		"/publish",
		m.UseHandle(h.Handle),
	)
}
