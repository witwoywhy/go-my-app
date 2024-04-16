package echo

import (
	"myapp/libs/app"

	"github.com/labstack/echo/v4"
)

func toEchoHandler(handler app.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return handler(echoCtxToWebFrameworkCtx(ctx))
	}
}
