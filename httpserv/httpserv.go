package httpserv

import (
	"myapp/libs/echo"
	"myapp/libs/gin"
)

func Run() {
	g := gin.New()
	g.UseMiddleware(gin.LogRequest())
	bindHelloHandler(g)
	bindPublishHandler(g)
	go g.ListenAndServe(":9998")

	e := echo.New()
	e.UseMiddleware(echo.LogRequest())
	bindHelloHandler(e)
	bindPublishHandler(e)
	go e.ListenAndServe(":9999")
}
