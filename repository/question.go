package repository

import (
	"github.com/google/uuid"
	"github.com/ikura-hamu/questions/domain"
)

type QuestionRepository interface {
	CreateQuestion(question string) (domain.Question, error)
	GetQuestionById(id uuid.UUID) (domain.Question, error)
	GetQuestions(limit int, offset int, answered bool) (int, []domain.Question, error)
	GetAllQuestions() (int, []domain.Question, error)
	CreateAnswer(id uuid.UUID, answer string, userId string) (domain.Question, error)
}
