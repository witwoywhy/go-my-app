package gin

import (
	"bytes"
	"fmt"
	"io"
	"myapp/libs/app"
	"myapp/libs/log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func LogRequest() app.HandleFunc {
	return func(ctx app.WebFrameworkContext, l log.Logger) error {
		writer := ctx.GetWriter().(gin.ResponseWriter)
		w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: writer}
		ctx.SetWriter(w)

		request := ctx.GetRequest().(*http.Request)

		b, _ := io.ReadAll(request.Body)
		request.Body = io.NopCloser(bytes.NewBuffer(b))

		msg := fmt.Sprintf("http-request | %s | %s%s |", request.Method, request.Host, request.URL.Path)
		l.LogHttpRequest(msg, request.Header, b)

		ctx.Next()

		msg = fmt.Sprintf("http-response | %s%s |", request.Host, request.URL.Path)
		l.LogHttpRequest(msg, w.Header(), w.body.Bytes())

		return nil
	}
}
