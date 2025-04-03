package http_gin

import (
	"fmt"
	"net/http"
	"real-time-messaging/producer/config"

	di "real-time-messaging/producer/internal"
	"real-time-messaging/producer/internal/adapters/inbound/rest/v1/handlers"
	ws "real-time-messaging/producer/internal/domain/services/websocket"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
		router.buildSwaggerRoutes(v1)
		router.buildIndexRoutes(v1)
	}

	r.Run(fmt.Sprintf(":%d", config.HTTP.Port))

	return router
}

func (r *Router) buildSwaggerRoutes(rg *gin.RouterGroup) {
	swaggerRoutes := rg.Group("/swagger")
	{
		swaggerRoutes.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}

func (r *Router) buildIndexRoutes(rg *gin.RouterGroup) {
	indexRoutes := rg.Group("/")
	{
		indexRoutes.GET("/health", handlers.HealthCheck)
	}
}

func (router *Router) buildWebSocketRoutes(rg *gin.RouterGroup) {
	websocketService := ws.NewWebsocketService(
		ws.WithLogger(router.container.Logger),
		ws.WithUpgrader(router.upgrader),
	)

	webSocketHandler := handlers.NewWebsocketHandler(
		router.upgrader,
		websocketService,
		router.container.Logger,
	)

	websocketRoutes := rg.Group("/ws/")
	{
		websocketRoutes.GET("", webSocketHandler.WebsocketReceive)
	}
}
