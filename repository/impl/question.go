package impl

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ikura-hamu/questions/domain"
	"github.com/jmoiron/sqlx"
)

type Question struct {
	Id        uuid.UUID `db:"id"`
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

func (q *questionRepository) CreateQuestion(question string) error {
	id := uuid.New()
	_, err := q.db.Exec("INSERT INTO `question` (`id`, `question`) VALUES (?, ?)", id, question)
	if err != nil {
		return fmt.Errorf("failed to create question: %w", err)
	}
	return nil
}

func (q *questionRepository) GetQuestionById(id uuid.UUID) (domain.Question, error) {
	var question domain.Question
	err := q.db.Get(&question, "SELECT * FROM `questions` WHERE `id` = ?", id)
	if err != nil {
		return domain.Question{}, fmt.Errorf("failed to get question by id: %w", err)
	}

	return domain.NewQuestion(question.Id, question.Question, question.Answer, question.CreatedAt), nil
}

func (q *questionRepository) GetQuestions(limit int, offset int) ([]domain.Question, error) {
	var questions []Question
	query := "SELECT * FROM `questions` LIMIT ? OFFSET ?"
	err := q.db.Select(&questions, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get questions with limit: %w", err)
	}

	res := make([]domain.Question, 0, len(questions))
	for i := range questions {
		res = append(res, domain.NewQuestion(questions[i].Id, questions[i].Question, questions[i].Answer, questions[i].CreatedAt))
	}
	return res, nil
}

func (q *questionRepository) GetAllQuestions() (int, []domain.Question, error) {
	var questions []Question
	query := "SELECT * FROM `questions`"
	err := q.db.Select(&questions, query)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to get questions with limit: %w", err)
	}

	res := make([]domain.Question, 0, len(questions))
	for i := range questions {
		res = append(res, domain.NewQuestion(questions[i].Id, questions[i].Question, questions[i].Answer, questions[i].CreatedAt))
	}
	return len(res), res, nil
}

func (q *questionRepository) CreateAnswer(id uuid.UUID, answer string) error {
	_, err := q.db.Exec("UPDATE `questions` SET `answer` = ? WHERE `id` = ?", answer, id)
	if err != nil {
		return fmt.Errorf("failed to create answer: %w", err)
	}
	return nil
}
