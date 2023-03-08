package impl

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Question struct {
	Id        int       `db:"id"`
	Question  string    `db:"question"`
	Answer    string    `db:"answer"`
	CreatedAt time.Time `db:"created_at"`
}

type questionRepository struct {
	db *sqlx.DB
}

func NewQuestionRepository(db *sqlx.DB) *questionRepository {
	return &questionRepository{db: db}
}
