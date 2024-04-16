package echo

import (
	"myapp/libs/app"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Context struct {
	ctx echo.Context
}

func echoCtxToWebFrameworkCtx(ctx echo.Context) app.WebFrameworkContext {
	return &Context{
		ctx: ctx,
	}
}

func (c *Context) BindJSON(obj any) error {
	return c.ctx.Bind(obj)
}

func (c *Context) GetHeader(key string) string {
	return c.ctx.Request().Header.Get(key)
}

func (c *Context) SetHeader(key string, value string) {
	c.ctx.Request().Header.Set(key, value)
}

func (c *Context) GetRequest() any {
	return c.ctx.Request()
}

func (c *Context) JSON(code int, obj any) error {
	return c.ctx.JSON(code, obj)
}

func (c *Context) GetWriter() any {
	return c.ctx.Response().Writer
}

func (c *Context) SetWriter(writer any) {
	c.ctx.Response().Writer = writer.(http.ResponseWriter)
}

func (c *Context) Next() {
	panic("unimplemented")
}
