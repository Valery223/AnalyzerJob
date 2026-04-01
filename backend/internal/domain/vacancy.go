package domain

import (
	"context"
	"time"
)

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
	// GetByID(ctx context.Context, id string) (*Vacancy, error)
}

// VacancyRepository  - interface для   БД
type VacancyRepository interface {
	Store(ctx context.Context, vacancy *Vacancy) error
	// GetByID(ctx context.Context, id string) (*Vacancy, error)
}
