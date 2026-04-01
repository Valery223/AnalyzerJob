package usecase

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
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
func (u *vacancyUsecase) GenerateQuestions(ctx context.Context, id string) ([]string, error) {
	//  Достаем вакансию из БД
	vacancy, err := u.vacancyRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	//  Формируем промпт для ИИ
	prompt := fmt.Sprintf("Ты IT-рекрутер. Напиши ровно 5 вопросов для технического собеседования по этой вакансии. Никаких вступлений, только сами вопросы, каждый с новой строки, без цифр и дефисов в начале. Вакансия: %s", vacancy.Description)

	encodedPrompt := url.QueryEscape(prompt)

	//API Pollinations
	aiURL := "https://text.pollinations.ai/prompt/" + encodedPrompt + "?model=openai"

	resp, err := http.Get(aiURL)
	if err != nil {
		return nil, fmt.Errorf("ошибка сети при обращении к ИИ: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ИИ вернул ошибку, статус: %d", resp.StatusCode)
	}

	// Читаем ответ
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения ответа ИИ: %v", err)
	}

	aiResponseText := string(bodyBytes)

	//Парсим ответ
	rawQuestions := strings.Split(aiResponseText, "\n")

	var questions []string
	for _, q := range rawQuestions {
		cleaned := strings.TrimSpace(q)
		cleaned = strings.TrimPrefix(cleaned, "-")
		cleaned = strings.TrimPrefix(cleaned, "*")
		cleaned = strings.TrimSpace(cleaned)

		if len(cleaned) > 5 {
			questions = append(questions, cleaned)
		}
	}

	if len(questions) == 0 {
		questions = []string{
			"Расскажите о вашем релевантном опыте для этой позиции?",
			"С какими сложностями вы сталкивались в похожих задачах?",
		}
	}

	// Сохраняем
	vacancy.AIQuestions = questions
	err = u.vacancyRepo.Update(ctx, vacancy)

	return questions, err
}
