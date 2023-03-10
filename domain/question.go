package domain

import (
	"time"

	"github.com/google/uuid"
)

type Question struct {
	Id        uuid.UUID `json:"id"`
	Question  string    `json:"question"`
	Answer    string    `json:"answer"`
	Answerer  string    `json:"answerer"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewQuestion(id uuid.UUID, question string, answer string,answerer string, createdAt time.Time, updatedAt time.Time) Question {
	return Question{
		Id:        id,
		Question:  question,
		Answer:    answer,
		Answerer: answerer,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
