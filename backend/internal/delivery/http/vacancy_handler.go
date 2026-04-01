package http

import (
	"net/http"

	"github.com/Valery223/AnalyzerJob/backend/internal/domain"
	"github.com/gin-gonic/gin"
)

// VacancyHandler хранит интерфейс UseCase
type VacancyHandler struct {
	VacancyUC domain.VacancyUsecase
}

// RegisterRoutes привязывает хэндлеры к роутеру Gin
func NewVacancyHandler(r *gin.RouterGroup, us domain.VacancyUsecase) {
	handler := &VacancyHandler{
		VacancyUC: us,
	}

	// Роуты для вакансий
	r.POST("/", handler.Create)
}

func (h *VacancyHandler) Create(c *gin.Context) {
	var vacancy domain.Vacancy

	// Парсим json из запроса
	if err := c.ShouldBindJSON(&vacancy); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// (для теста)
	vacancy.UserID = "00000000-0000-0000-0000-000000000000"

	err := h.VacancyUC.Create(c.Request.Context(), &vacancy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, vacancy)
}
