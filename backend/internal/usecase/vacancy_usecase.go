package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/Valery223/AnalyzerJob/backend/internal/domain"
)

type vacancyUsecase struct {
	vacancyRepo domain.VacancyRepository
}

// NewVacancyUsecase - конструктор
func NewVacancyUsecase(v domain.VacancyRepository) domain.VacancyUsecase {
	return &vacancyUsecase{
		vacancyRepo: v,
	}
}

func (u *vacancyUsecase) Create(ctx context.Context, vacancy *domain.Vacancy) error {
	// Здесь  бизнес-логика(валидация):
	if vacancy.Title == "" {
		return errors.New("title is required")
	}

	vacancy.CreatedAt = time.Now()

	// Передаем сохранение в слой БД
	return u.vacancyRepo.Store(ctx, vacancy)
}

func (u *vacancyUsecase) GetByID(ctx context.Context, id string) (*domain.Vacancy, error) {
	// TODO
	return u.vacancyRepo.GetByID(ctx, id)
}

func (u *vacancyUsecase) Fetch(ctx context.Context, userID string, filter domain.VacancyFilter) ([]*domain.Vacancy, error) {
	// TODO
	return u.vacancyRepo.Fetch(ctx, userID, filter)
}

func (u *vacancyUsecase) Delete(ctx context.Context, id string) error {
	// TODO
	return u.vacancyRepo.Delete(ctx, id)
}
