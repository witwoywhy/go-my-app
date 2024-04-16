package gin

import (
	"myapp/libs/app"
	"myapp/libs/log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type App struct {
	gin *gin.Engine
}

func New() app.App {
	return &App{
		gin: gin.New(),
	}
}

func (a *App) Register(method string, relativePath string, handlers ...app.HandlerFunc) {
	h := toGinHandler(handlers...)
	switch method {
	case http.MethodGet:
		a.gin.GET(relativePath, h...)
	case http.MethodPost:
		a.gin.POST(relativePath, h...)
	}
}

func (a *App) UseMiddleware(handles ...any) {
	for _, handle := range handles {
		h := toGinHandler(handle.(app.HandlerFunc))
		a.gin.Use(h...)
	}
}

func (a *App) ListenAndServe(addr string) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := http.ListenAndServe(addr, a.gin); err != nil {
			log.L.Fatal(err)
		}
	}()
	wg.Wait()
}
