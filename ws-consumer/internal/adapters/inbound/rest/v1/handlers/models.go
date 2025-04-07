package handlers

type LoginRequest struct {
	Email    string `json:"email" binding:"required" example:"john.doe@example.com"`
	Password string `json:"password" binding:"required" example:"password123"`
}

type LoginResponse struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}
