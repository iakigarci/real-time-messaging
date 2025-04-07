package http_gin

import (
	"fmt"
	"real-time-messaging/consumer/config"
	di "real-time-messaging/consumer/internal"
	"real-time-messaging/consumer/internal/adapters/inbound/rest/v1/handlers"
	"real-time-messaging/consumer/internal/domain/services/auth"
	"real-time-messaging/consumer/internal/domain/services/consumer"
	"real-time-messaging/consumer/internal/domain/services/user"

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
		router.buildAuthRoutes(v1)
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

func (r *Router) buildWebSocketRoutes(rg *gin.RouterGroup) {
	consumerService := consumer.New(
		consumer.WithLogger(r.container.Logger),
		consumer.WithWebsocket(r.container.WebsocketPort),
		consumer.WithMessageProducer(r.container.MessageProducer),
	)

	webSocketHandler := handlers.NewWebsocketHandler(
		consumerService,
		r.container.Logger,
	)

	websocketRoutes := rg.Group("/ws/", AuthMiddleware(r.container.AuthPort))
	{
		websocketRoutes.GET("", webSocketHandler.WebsocketReceive)
	}
}

func (r *Router) buildAuthRoutes(rg *gin.RouterGroup) {
	userService := user.New(
		user.WithUserRepository(r.container.UserRepository),
		user.WithLogger(r.container.Logger),
	)

	authService := auth.New(
		auth.WithAuthService(r.container.AuthPort),
		auth.WithLogger(r.container.Logger),
	)

	authHandler := handlers.NewAuthHandler(authService, userService)

	authRoutes := rg.Group("/auth")
	{
		authRoutes.POST("/login", authHandler.Login)
	}
}
