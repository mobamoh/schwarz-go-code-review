package dep

import (
	"github.com/gin-gonic/gin"
	"github.com/mobamoh/schwarz-go-code-review/coupon-service/pkg/transport/bind"
	"net"
	"net/http"
)

type App struct {
	listener net.Listener
	server   *http.Server
}

func (a *App) Start() error {
	return a.server.Serve(a.listener)
}

func NewListener() (net.Listener, error) {
	return net.Listen("tcp", ":8080")
}

func NewRouter(handler bind.CouponHandler) http.Handler {
	router := gin.Default()
	group := router.Group("/api/coupon")
	group.POST("/apply", handler.HandleApply)
	group.POST("/", handler.HandleCreate)
	group.GET("/", handler.HandleList)
	return router
}

func NewApp(listener net.Listener, handler http.Handler) *App {
	return &App{
		listener: listener,
		server:   &http.Server{Handler: handler},
	}
}
