package domain

import "context"

type User struct {
	ID           string `json:"id"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
}

// Запросы от фронтенда
type AuthRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type AuthUsecase interface {
	Register(ctx context.Context, req AuthRequest) error
	Login(ctx context.Context, req AuthRequest) (string, error) // Возвращает JWT токен
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetByEmail(ctx context.Context, email string) (*User, error)
}
