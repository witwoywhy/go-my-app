package app

import (
	"myapp/libs/log"

	"github.com/google/uuid"
)

type App interface {
	Register(method, relativePath string, handlers ...HandlerFunc)
	UseMiddleware(handles ...any)
	ListenAndServe(addr string)
}

const (
	traceID = "traceId"
	spanID  = "spanId"
)

func GetIDByKey(key string, ctx WebFrameworkContext) string {
	id := ctx.GetHeader(key)
	if id == "" {
		id = uuid.NewString()
		ctx.SetHeader(key, id)
	}

	return id
}

func NewLogFromCtx(ctx WebFrameworkContext) log.Logger {
	return &log.Log{
		TraceID: GetIDByKey(traceID, ctx),
		SpanID:  GetIDByKey(spanID, ctx),
	}
}

func UseHandle(handle HandleFunc) HandlerFunc {
	return func(ctx WebFrameworkContext) error {
		l := NewLogFromCtx(ctx)
		return handle(ctx, l)
	}
}
