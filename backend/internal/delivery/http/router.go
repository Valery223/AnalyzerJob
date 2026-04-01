package http

import (
	"net/http"

	_ "github.com/Valery223/AnalyzerJob/backend/docs"
	"github.com/Valery223/AnalyzerJob/backend/internal/domain"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRouter собирает все роуты приложения
func SetupRouter(vacancyUC domain.VacancyUsecase, authUC domain.AuthUsecase) *gin.Engine {
	r := gin.Default()

	// Отключаем перенаправление trailing slash
	// r.RedirectTrailingSlash = false

	// Подключаем cors
	r.Use(corsMiddleware())

	//ping
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// api v1
	v1 := r.Group("/api/v1")
	{
		// Регистрация роутов авторизации
		authGroup := v1.Group("/auth")
		RegisterAuthRoutes(authGroup, authUC)

		// Роуты для вакансий
		vacancyGroup := v1.Group("/vacancies")
		vacancyGroup.Use(AuthMiddleware()) // Все роуты вакансий требуют авторизации
		RegisterVacancyRoutes(vacancyGroup, vacancyUC)

		// другие, типо auth будут
	}

	return r
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173") // Адрес React
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
