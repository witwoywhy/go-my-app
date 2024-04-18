package gin

import (
	"myapp/libs/app"

	"github.com/gin-gonic/gin"
)

func toGinHandler(handlers ...app.HandleFunc) []gin.HandlerFunc {
	return []gin.HandlerFunc{
		func(ctx *gin.Context) {
			for _, handler := range handlers {
				c := ginCtxToWebFrameworkCtx(ctx)
				l := app.NewLogFromCtx(c)
				handler(c, l)
			}
		},
	}
}
