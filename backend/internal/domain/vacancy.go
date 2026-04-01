package domain

import (
	"context"
	"time"
)

// VacancyFilter - параметры для поиска
type VacancyFilter struct {
	SearchQuery string // Поиск по названию или компании
}

// Vacancy
type Vacancy struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	Title       string    `json:"title"`
	Company     string    `json:"company"`
	Description string    `json:"description"`
	AIQuestions []string  `json:"ai_questions"`
	CreatedAt   time.Time `json:"created_at"`
}

// VacancyUsecase - interface для  Usecase
type VacancyUsecase interface {
	Create(ctx context.Context, vacancy *Vacancy) error
	GetByID(ctx context.Context, id string) (*Vacancy, error)
	Fetch(ctx context.Context, userID string, filter VacancyFilter) ([]*Vacancy, error)
	Delete(ctx context.Context, id string) error
}

// VacancyRepository  - interface для   БД
type VacancyRepository interface {
	Store(ctx context.Context, vacancy *Vacancy) error
	GetByID(ctx context.Context, id string) (*Vacancy, error)
	Fetch(ctx context.Context, userID string, filter VacancyFilter) ([]*Vacancy, error)
	Delete(ctx context.Context, id string) error
}
