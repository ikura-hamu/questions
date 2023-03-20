package impl

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ikura-hamu/questions/domain"
	"github.com/ikura-hamu/questions/repository"
	"github.com/jmoiron/sqlx"
)

type Question struct {
	Id        uuid.UUID      `db:"id"`
	Question  string         `db:"question"`
	Answer    sql.NullString `db:"answer"`
	Answerer  sql.NullString `db:"answerer"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt time.Time      `db:"updated_at"`
}

type questionRepository struct {
	db *sqlx.DB
}

func NewQuestionRepository(db *sqlx.DB) repository.QuestionRepository {
	return &questionRepository{
		db: db,
	}
}

func (q *questionRepository) CreateQuestion(question string) (domain.Question, error) {
	id := uuid.New()
	_, err := q.db.Exec("INSERT INTO `questions` (`id`, `question`) VALUES (?, ?)", id, question)
	if err != nil {
		return domain.Question{}, fmt.Errorf("failed to create question: %w", err)
	}
	return domain.NewQuestion(id, question, "", "", time.Now(), time.Now()), nil
}

func (q *questionRepository) GetQuestionById(id uuid.UUID) (domain.Question, error) {
	var question Question
	err := q.db.Get(&question, "SELECT * FROM `questions` WHERE `id` = ?", id)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.Question{}, repository.ErrNoRecord
	}
	if err != nil {
		return domain.Question{}, fmt.Errorf("failed to get question by id: %w", err)
	}

	return domain.NewQuestion(question.Id, question.Question, question.Answer.String, question.Answerer.String, question.CreatedAt, question.UpdatedAt), nil
}

func (q *questionRepository) GetQuestions(limit int, offset int) (int, []domain.Question, error) {
	var count int
	err := q.db.Get(&count, "SELECT COUNT(*) FROM `questions`")
	if err != nil {
		return 0, nil, fmt.Errorf("failed to get questions count: %w", err)
	}

	var questions []Question
	query := "SELECT * FROM `questions` ORDER BY `created_at` DESC LIMIT ? OFFSET ?"
	err = q.db.Select(&questions, query, limit, offset)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to get questions with limit: %w", err)
	}

	res := make([]domain.Question, 0, len(questions))
	for _, question := range questions {
		res = append(res, domain.NewQuestion(
			question.Id,
			question.Question,
			question.Answer.String,
			question.Answerer.String,
			question.CreatedAt,
			question.UpdatedAt))
	}
	return count, res, nil
}

func (q *questionRepository) GetAllQuestions() (int, []domain.Question, error) {
	var questions []Question
	query := "SELECT * FROM `questions`"
	err := q.db.Select(&questions, query)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to get questions with limit: %w", err)
	}

	res := make([]domain.Question, 0, len(questions))
	for _, question := range questions {
		res = append(res, domain.NewQuestion(
			question.Id,
			question.Question,
			question.Answer.String,
			question.Answerer.String,
			question.CreatedAt,
			question.UpdatedAt))
	}
	return len(res), res, nil
}

func (q *questionRepository) CreateAnswer(id uuid.UUID, answer string, answerer uuid.UUID) (domain.Question, error) {
	var question Question
	err := q.db.Get(&question, "SELECT * FROM `questions` WHERE `id` = ?", id)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.Question{}, repository.ErrNoRecord
	}
	if err != nil {
		return domain.Question{}, err
	}

	updatedAt := time.Now()

	_, err = q.db.Exec("UPDATE `questions` SET `answer` = ?,`answerer` = ?, `updated_at` = ? WHERE `id` = ?", answer, answerer, updatedAt, id)
	if err != nil {
		return domain.Question{}, fmt.Errorf("failed to create answer: %w", err)
	}
	return domain.NewQuestion(id, question.Question, answer, answerer.String(), question.CreatedAt, updatedAt), nil
}
