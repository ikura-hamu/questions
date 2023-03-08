package domain

import (
	"time"

	"github.com/google/uuid"
)

type Question struct {
	Id        uuid.UUID
	Question  string
	Answer    string
	CreatedAt time.Time
}

func NewQuestion(id uuid.UUID, question string, answer string, createdAt time.Time) Question {
	return Question{
		Id:        id,
		Question:  question,
		Answer:    answer,
		CreatedAt: createdAt,
	}
}
