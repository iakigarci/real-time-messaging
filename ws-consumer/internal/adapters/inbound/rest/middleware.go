package http_gin

import (
	"net/http"
	port "real-time-messaging/consumer/internal/domain/ports"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With", "Sec-WebSocket-Key", "Sec-WebSocket-Version", "Sec-WebSocket-Extensions", "Sec-WebSocket-Protocol", "Connection", "Upgrade"},
		ExposeHeaders:    []string{"Content-Length", "Upgrade", "Connection", "Sec-WebSocket-Accept"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}

func AuthMiddleware(authClient port.Authentication) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing authorization header"})
			return
		}

		token = strings.TrimPrefix(token, "Bearer ")

		userID, err := authClient.ValidateToken(c.Request.Context(), token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		c.Set("user_id", userID)
		c.Next()
	}
}
