package repository

import (
	"github.com/ikura-hamu/questions/domain"
)

type QuestionRepository interface {
	CreateQuestion(question string) (domain.Question, error)
	GetQuestionById(id int) (domain.Question, error)
	GetQuestions(limit int, offset int) ([]domain.Question, error)
	CreateAnswer(answer string) (domain.Question, error)
}
