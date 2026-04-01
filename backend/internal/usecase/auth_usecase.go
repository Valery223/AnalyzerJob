package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/Valery223/AnalyzerJob/backend/internal/domain"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Ключ для подписи токенов (в проде будет из .env)
var jwtSecret = []byte("secret_key")

type authUsecase struct {
	userRepo domain.UserRepository
}

func NewAuthUsecase(ur domain.UserRepository) domain.AuthUsecase {
	return &authUsecase{userRepo: ur}
}

func (u *authUsecase) Register(ctx context.Context, req domain.AuthRequest) error {
	// 1. Хэшируем пароль
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// 2. Создаем юзера и передаем в БД
	user := &domain.User{
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
	}

	return u.userRepo.Create(ctx, user)
}

func (u *authUsecase) Login(ctx context.Context, req domain.AuthRequest) (string, error) {
	// 1. Ищем пользователя по email
	user, err := u.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	// 2. Сравниваем хэш из БД с паролем от фронтенда
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	// 3. Генерируем JWT токен
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,                               // Зашиваем ID пользователя в токен
		"exp":     time.Now().Add(time.Hour * 72).Unix(), // Токен живет 3 дня
	})

	// Подписываем токен ключом
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
