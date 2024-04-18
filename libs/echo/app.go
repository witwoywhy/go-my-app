package echo

import (
	"myapp/libs/app"
	"myapp/libs/log"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
)

type App struct {
	echo *echo.Echo
}

func New() app.App {
	return &App{
		echo: echo.New(),
	}
}

func (a *App) Register(method string, relativePath string, handlers ...app.HandleFunc) {
	switch method {
	case http.MethodGet:
		a.echo.GET(relativePath, toEchoHandler(handlers[0]))
	case http.MethodPost:
		a.echo.POST(relativePath, toEchoHandler(handlers[0]))
	}
}

func (a *App) ListenAndServe(addr string) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := http.ListenAndServe(addr, a.echo); err != nil {
			log.L.Fatal(err)
		}
	}()
	wg.Wait()
}

func (a *App) UseMiddleware(fn ...any) {
	for _, f := range fn {
		a.echo.Use(f.(echo.MiddlewareFunc))
	}
}
