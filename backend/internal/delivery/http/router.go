package http

import (
	"net/http"

	"github.com/Valery223/AnalyzerJob/backend/internal/domain"
	"github.com/gin-gonic/gin"
)

// SetupRouter собирает все роуты приложения
func SetupRouter(vacancyUC domain.VacancyUsecase) *gin.Engine {
	r := gin.Default()

	// Подключаем cors
	r.Use(corsMiddleware())

	//ping
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// api v1
	v1 := r.Group("/api/v1")
	{
		vacancyGroup := v1.Group("/vacancies")
		RegisterVacancyRoutes(vacancyGroup, vacancyUC)

		// другие, типо auth будут
	}

	return r
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
	}
}
