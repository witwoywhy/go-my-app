package echo

import (
	"bytes"
	"fmt"
	"io"
	"myapp/libs/app"
	"net/http"

	"github.com/labstack/echo/v4"
)

type responseBodyWriter struct {
	http.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func LogRequest() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := echoCtxToWebFrameworkCtx(c)
			l := app.NewLogFromCtx(ctx)

			writer := ctx.GetWriter().(http.ResponseWriter)
			w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: writer}
			ctx.SetWriter(w)

			request := ctx.GetRequest().(*http.Request)
			b, _ := io.ReadAll(request.Body)
			request.Body = io.NopCloser(bytes.NewBuffer(b))

			msg := fmt.Sprintf("http-request | %s | %s%s |", request.Method, request.Host, request.URL.Path)
			l.LogHttpRequest(msg, request.Header, b)

			err := next(c)

			msg = fmt.Sprintf("http-response | %s%s |", request.Host, request.URL.Path)
			l.LogHttpRequest(msg, w.Header(), w.body.Bytes())

			c.Response()
			return err
		}
	}
}
