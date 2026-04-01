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
	r.POST("/:id/generate", handler.Generate)
}

// @Summary Создать новую вакансию
// @Description Создает вакансию для текущего пользователя
// @Tags vacancies
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT токен"
// @Param input body domain.Vacancy true "Данные вакансии"
// @Success 201 {object} domain.Vacancy
// @Failure 400 {object} map[string]string
// @Router /api/v1/vacancies [post]
func (h *VacancyHandler) Create(c *gin.Context) {
	var vacancy domain.Vacancy

	// Парсим json из запроса
	if err := c.ShouldBindJSON(&vacancy); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error parse": err.Error()})
		return
	}

	// Достаем user_id из контекста, который положил AuthMiddleware
	userID := c.GetString("user_id")
	vacancy.UserID = userID

	err := h.VacancyUC.Create(c.Request.Context(), &vacancy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, vacancy)
}

// @Summary Получить вакансию по ID
// @Description Получает полную информацию о вакансии по идентификатору
// @Tags vacancies
// @Produce json
// @Param Authorization header string true "JWT токен"
// @Param id path string true "Идентификатор вакансии"
// @Success 200 {object} domain.Vacancy
// @Failure 404 {object} map[string]string "error"
// @Router /api/v1/vacancies/{id} [get]
func (h *VacancyHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	vacancy, err := h.VacancyUC.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vacancy not found"})
		return
	}

	c.JSON(http.StatusOK, vacancy)
}

// @Summary Получить все вакансии пользователя
// @Description Получает список вакансий текущего пользователя с опциональным поиском
// @Tags vacancies
// @Produce json
// @Param Authorization header string true "JWT токен"
// @Param search query string false "Поисковый запрос по названию или компании"
// @Success 200 {array} domain.Vacancy
// @Failure 500 {object} map[string]string "error"
// @Router /api/v1/vacancies [get]
func (h *VacancyHandler) Fetch(c *gin.Context) {
	// Достаем ?search= из URL
	filter := domain.VacancyFilter{
		SearchQuery: c.Query("search"),
	}

	userID := c.GetString("user_id")

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

// @Summary Удалить вакансию
// @Description Удаляет вакансию по идентификатору
// @Tags vacancies
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT токен"
// @Param id path string true "Идентификатор вакансии"
// @Success 200 {object} map[string]string "message"
// @Failure 500 {object} map[string]string "error"
// @Router /api/v1/vacancies/{id} [delete]
func (h *VacancyHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	err := h.VacancyUC.Delete(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully deleted"})
}

// @Summary Генерировать вопросы для вакансии
// @Description Использует AI для генерации вопросов на основе текста вакансии
// @Tags vacancies
// @Produce json
// @Param Authorization header string true "JWT токен"
// @Param id path string true "Идентификатор вакансии"
// @Success 200 {object} map[string][]string "questions"
// @Failure 500 {object} map[string]string "error"
// @Router /api/v1/vacancies/{id}/generate [post]
func (h *VacancyHandler) Generate(c *gin.Context) {
	id := c.Param("id")

	// Вызываем логику генерации
	questions, err := h.VacancyUC.GenerateQuestions(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate questions"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"questions": questions})
}
