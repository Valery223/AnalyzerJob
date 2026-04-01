package postgresrep

import (
	"context"
	"fmt"

	"github.com/Valery223/AnalyzerJob/backend/internal/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type vacancyRepository struct {
	db *pgxpool.Pool
}

// NewVacancyRepository - конструктор
func NewVacancyRepository(db *pgxpool.Pool) domain.VacancyRepository {
	return &vacancyRepository{db: db}
}

func (r *vacancyRepository) Store(ctx context.Context, v *domain.Vacancy) error {
	query := `
		INSERT INTO vacancies (user_id, title, company, description, created_at) 
		VALUES ($1, $2, $3, $4, $5) 
		RETURNING id`

	// Выполняем запрос и сразу записываем сгенерированный базой UUID обратно в структуру
	err := r.db.QueryRow(ctx, query, v.UserID, v.Title, v.Company, v.Description, v.CreatedAt).Scan(&v.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *vacancyRepository) GetByID(ctx context.Context, id string) (*domain.Vacancy, error) {
	query := `SELECT id, user_id, title, company, description, ai_questions, created_at FROM vacancies WHERE id = $1`
	v := &domain.Vacancy{}

	err := r.db.QueryRow(ctx, query, id).Scan(
		&v.ID, &v.UserID, &v.Title, &v.Company, &v.Description, &v.AIQuestions, &v.CreatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("vacancy not found")
		}
		return nil, err
	}
	return v, nil
}

// Fetch - получить список с учетом фильтров
func (r *vacancyRepository) Fetch(ctx context.Context, userID string, filter domain.VacancyFilter) ([]*domain.Vacancy, error) {
	// Базовый запрос
	query := `SELECT id, user_id, title, company, description, ai_questions, created_at FROM vacancies WHERE user_id = $1`
	args := []interface{}{userID}
	argCount := 2

	// Если есть строка поиска, добавляем условие ILIKE
	if filter.SearchQuery != "" {
		query += fmt.Sprintf(` AND (title ILIKE $%d OR company ILIKE $%d)`, argCount, argCount)
		args = append(args, "%"+filter.SearchQuery+"%")
	}

	// Сортируем от новых к старым
	query += ` ORDER BY created_at DESC`

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*domain.Vacancy
	for rows.Next() {
		v := &domain.Vacancy{}
		err := rows.Scan(&v.ID, &v.UserID, &v.Title, &v.Company, &v.Description, &v.AIQuestions, &v.CreatedAt)
		if err != nil {
			return nil, err
		}
		results = append(results, v)
	}

	return results, nil
}

// Delete - удалить вакансию
func (r *vacancyRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM vacancies WHERE id = $1`
	cmdTag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("vacancy not found")
	}
	return nil
}
