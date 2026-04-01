package http

import (
	"net/http"

	"github.com/Valery223/AnalyzerJob/backend/internal/domain"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	AuthUC domain.AuthUsecase
}

func RegisterAuthRoutes(r *gin.RouterGroup, uc domain.AuthUsecase) {
	handler := &AuthHandler{AuthUC: uc}

	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req domain.AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid format. Need email and password (min 6 chars)"})
		return
	}

	if err := h.AuthUC.Register(c.Request.Context(), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not register user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User successfully registered!"})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req domain.AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	token, err := h.AuthUC.Login(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Отдаем токен фронту
	c.JSON(http.StatusOK, gin.H{"token": token})
}
