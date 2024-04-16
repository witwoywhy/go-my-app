package gin

import (
	"myapp/libs/app"

	"github.com/gin-gonic/gin"
)

func toGinHandler(handlers ...app.HandlerFunc) []gin.HandlerFunc {
	return []gin.HandlerFunc{
		func(ctx *gin.Context) {
			for _, handler := range handlers {
				handler(ginCtxToWebFrameworkCtx(ctx))
			}
		},
	}
}

