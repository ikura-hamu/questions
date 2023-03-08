package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/ikura-hamu/questions/domain"
)

type QuestionRepository interface {
	CreateQuestion(question string) (domain.Question, error)
	GetQuestionById(id uuid.UUID) (domain.Question, error)
	GetQuestions(limit int, offset int) ([]domain.Question, error)
	GetAllQuestions() (int, []domain.Question, error)
	CreateAnswer(id uuid.UUID, answer string) (domain.Question, error)
}

var ErrNoRecord = errors.New("no record")
