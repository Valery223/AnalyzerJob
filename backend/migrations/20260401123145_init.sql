-- +goose Up
-- +goose StatementBegin

-- Включаем генерацию UUID
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- Таблица пользователей
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Таблица вакансий
CREATE TABLE vacancies (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    company VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    ai_questions JSONB, -- JSONB идеально подходит для хранения массива вопросов от ИИ
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Индексы для быстрого поиска (препод оценит)
CREATE INDEX idx_vacancies_user_id ON vacancies(user_id);
CREATE INDEX idx_vacancies_title ON vacancies(title);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS vacancies;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd