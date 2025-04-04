package http_gin

import (
	"fmt"
	"real-time-messaging/consumer/config"
	di "real-time-messaging/consumer/internal"
	"real-time-messaging/consumer/internal/adapters/inbound/rest/v1/handlers"
	"real-time-messaging/consumer/internal/domain/services/consumer"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	Router    *gin.Engine
	container *di.Container
}

func New(config *config.Config, container *di.Container) *Router {
	r := gin.Default()

	router := &Router{
		Router:    r,
		container: container,
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
	consumerService := consumer.NewConsumerService(
		consumer.WithLogger(router.container.Logger),
		consumer.WithWebsocket(router.container.WebsocketPort),
	)

	webSocketHandler := handlers.NewWebsocketHandler(
		consumerService,
		router.container.Logger,
	)

	websocketRoutes := rg.Group("/ws/")
	{
		websocketRoutes.GET("", webSocketHandler.WebsocketReceive)
	}
}
