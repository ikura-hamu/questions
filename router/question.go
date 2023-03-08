package router

import (
	"github.com/ikura-hamu/questions/repository"
	"github.com/labstack/echo/v4"
)

type QuestionHandler interface {
	PostQuestionHandler(c echo.Context) error
	GetQuestionByIdHandler(c echo.Context) error
	GetQuestions(c echo.Context) error
	PostAnswerHandler(c echo.Context) error
}

type questionHandler struct {
	r *repository.QuestionRepository
}

func NewQuestionHandler(r *repository.QuestionRepository) *questionHandler {
	return &questionHandler{r: r}
}
