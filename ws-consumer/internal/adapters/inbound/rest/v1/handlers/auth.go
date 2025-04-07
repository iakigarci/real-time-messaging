package handlers

import (
	port "real-time-messaging/consumer/internal/domain/ports"
	httpserver "real-time-messaging/consumer/pkg/http"

	"github.com/gin-gonic/gin"
)

// @Summary User login
// @Description Authenticate user and return JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param login body LoginRequest true "Login credentials"
// @Success 200 {object} LoginResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /v1/auth/login [post]

type AuthHandler struct {
	authService port.AuthService
	userService port.UserService
}

func NewAuthHandler(as port.AuthService, us port.UserService) *AuthHandler {
	return &AuthHandler{
		authService: as,
		userService: us,
	}
}

// Login godoc
// @Summary User login
// @Description Authenticate user and return JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param login body LoginRequest true "Login credentials"
// @Success 200 {object} httpserver.SuccessResponseData
// @Failure 400 {object} httpserver.ErrorResponseData
// @Failure 401 {object} httpserver.ErrorResponseData
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		httpserver.ErrorResponse(c, err)
		return
	}

	user, err := h.userService.GetUserByCredentials(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		httpserver.ErrorResponse(c, err)
		return
	}

	token, err := h.authService.GenerateToken(c.Request.Context(), user)
	if err != nil {
		httpserver.ErrorResponse(c, err)
		return
	}

	httpserver.SuccessResponse(c, LoginResponse{Token: token})
}
