package echo

import (
	"myapp/libs/app"

	"github.com/labstack/echo/v4"
)

func toEchoHandler(handler app.HandleFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		c := echoCtxToWebFrameworkCtx(ctx)
		l := app.NewLogFromCtx(c)
		return handler(c, l)
	}
}
