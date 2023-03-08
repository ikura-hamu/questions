package repository

import (
	"github.com/google/uuid"
	"github.com/ikura-hamu/questions/domain"
)

type QuestionRepository interface {
	CreateQuestion(question string) error
	GetQuestionById(id uuid.UUID) (domain.Question, error)
	GetQuestions(limit int, offset int) ([]domain.Question, error)
	GetAllQuestions() (int, []domain.Question, error)
	CreateAnswer(answer string) error
}
