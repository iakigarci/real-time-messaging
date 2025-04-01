package http_gin

import (
	"fmt"
	"real-time-messaging/consumer/config"
	di "real-time-messaging/consumer/internal"

	"github.com/gin-gonic/gin"
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

	r.Run(fmt.Sprintf(":%d", config.HTTP.Port))

	return router
}
