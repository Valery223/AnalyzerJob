package postgresrep

import (
	"context"

	"github.com/Valery223/AnalyzerJob/backend/internal/domain"
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
