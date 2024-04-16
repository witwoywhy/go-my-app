package gin

import (
	"myapp/libs/app"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Context struct {
	ctx *gin.Context
}

func ginCtxToWebFrameworkCtx(ginCtx *gin.Context) app.WebFrameworkContext {
	return &Context{
		ctx: ginCtx,
	}
}

func (c *Context) BindJSON(obj any) error {
	return c.ctx.ShouldBindBodyWith(obj, binding.JSON)
}

func (c *Context) JSON(code int, obj any) error {
	c.ctx.JSON(code, obj)
	return nil
}

func (c *Context) GetHeader(key string) string {
	return c.ctx.GetHeader(key)
}

func (c *Context) SetHeader(key string, value string) {
	c.ctx.Request.Header.Add(key, value)
}

func (c *Context) Next() {
	c.ctx.Next()
}

func (c *Context) Skip() {
	c.ctx.Next()
}

func (c *Context) GetRequest() any {
	return c.ctx.Request
}

func (c *Context) GetWriter() any {
	return c.ctx.Writer
}

func (c *Context) SetWriter(writer any) {
	c.ctx.Writer = writer.(gin.ResponseWriter)
}
