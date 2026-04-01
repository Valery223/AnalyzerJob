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
func RegisterVacancyRoutes(r *gin.RouterGroup, uc domain.VacancyUsecase) {
	handler := &VacancyHandler{VacancyUC: uc}

	r.POST("/", handler.Create)
	r.GET("/", handler.Fetch)
	r.GET("/:id", handler.GetByID)
	r.DELETE("/:id", handler.Delete)
}

func (h *VacancyHandler) Create(c *gin.Context) {
	var vacancy domain.Vacancy

	// Парсим json из запроса
	if err := c.ShouldBindJSON(&vacancy); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error parse": err.Error()})
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

func (h *VacancyHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	vacancy, err := h.VacancyUC.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vacancy not found"})
		return
	}

	c.JSON(http.StatusOK, vacancy)
}

func (h *VacancyHandler) Fetch(c *gin.Context) {
	// Достаем ?search= из URL
	filter := domain.VacancyFilter{
		SearchQuery: c.Query("search"),
	}

	// Хардкод
	userID := "00000000-0000-0000-0000-000000000000"

	vacancies, err := h.VacancyUC.Fetch(c.Request.Context(), userID, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Если вакансий нет, отдаем пустой массив
	if vacancies == nil {
		vacancies = []*domain.Vacancy{}
	}

	c.JSON(http.StatusOK, vacancies)
}

func (h *VacancyHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	err := h.VacancyUC.Delete(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully deleted"})
}
