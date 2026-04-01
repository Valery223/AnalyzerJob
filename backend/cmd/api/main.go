package main

import (
	"context"
	"log"

	"github.com/Valery223/AnalyzerJob/backend/internal/delivery/http"
	postgresrep "github.com/Valery223/AnalyzerJob/backend/internal/repository/postgresRep"
	"github.com/Valery223/AnalyzerJob/backend/internal/usecase"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	// 1. Подключение к БД PostgreSQL
	dsn := "host=localhost port=5432 user=postgres password=password dbname=ai_jobs sslmode=disable"
	dbPool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer dbPool.Close()

	// 2. Инициализация слоев
	// Репозиторий
	vacancyRepo := postgresrep.NewVacancyRepository(dbPool)
	// Бизнес-логика
	vacancyUseCase := usecase.NewVacancyUsecase(vacancyRepo)

	// 3. Настройка Роутера
	router := http.SetupRouter(vacancyUseCase)

	// 4. Запуск сервера
	log.Println("Server is running on port 8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
