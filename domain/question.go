package domain

import (
	"time"

	"github.com/google/uuid"
)

type Question struct {
	Id        uuid.UUID `json:"id"`
	Question  string    `json:"question"`
	Answer    string    `json:"answer"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewQuestion(id uuid.UUID, question string, answer string, createdAt time.Time) Question {
	return Question{
		Id:        id,
		Question:  question,
		Answer:    answer,
		CreatedAt: createdAt,
	}
}
