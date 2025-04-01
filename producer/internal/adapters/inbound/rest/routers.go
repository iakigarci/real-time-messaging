package http_gin

import (
	"fmt"
	"net/http"
	"real-time-messaging/producer/config"

	di "real-time-messaging/producer/internal"
	"real-time-messaging/producer/internal/adapters/inbound/rest/v1/handlers"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Router struct {
	Router    *gin.Engine
	container *di.Container
	upgrader  websocket.Upgrader
}

func New(config *config.Config, container *di.Container) *Router {
	r := gin.Default()
	router := &Router{
		Router:    r,
		container: container,
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}

	r.SetTrustedProxies([]string{"127.0.0.1", "::1"})

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(CORSMiddleware())

	v1 := r.Group("/v1")
	{
		router.buildWebSocketRoutes(v1)
	}

	r.Run(fmt.Sprintf(":%d", config.HTTP.Port))

	return router
}

func (router *Router) buildWebSocketRoutes(v1 *gin.RouterGroup) {
	webSocketHandler := handlers.NewWebsocketHandler()
	v1.GET("/ws", webSocketHandler.Websocket)
}
