package router

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/ikura-hamu/questions/repository"
	traq "github.com/ikura-hamu/questions/traQ"
	"github.com/labstack/echo/v4"
)

type QuestionHandler interface {
	//POST /question
	PostQuestionHandler(c echo.Context) error
	//GET /question/:questionId
	GetQuestionByIdHandler(c echo.Context) error
	//GET /question?limit=10&offset=0
	GetQuestionsHandler(c echo.Context) error
	//POST /question/:questionId/answer
	PostAnswerHandler(c echo.Context) error
}

type questionHandler struct {
	r repository.QuestionRepository
}

func NewQuestionHandler(r repository.QuestionRepository) QuestionHandler {
	return &questionHandler{r: r}
}

type PostQuestionRequest struct {
	Question string `json:"question"`
}

type PostQuestionResponse struct {
	Id        uuid.UUID `json:"id"`
	Question  string    `json:"question"`
	CreatedAt time.Time `json:"created_at"`
}

func (h *questionHandler) PostQuestionHandler(c echo.Context) error {
	var req PostQuestionRequest
	err := c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "bad request body")
	}

	question, err := h.r.CreateQuestion(req.Question)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("failed to create question: %w", err))
	}

	lines := regexp.MustCompile("\n|\r\n").Split(question.Question, -1)
	q := ""
	for i := range lines {
		q = fmt.Sprintf("%v> %v\n", q, lines[i])
	}

	message := fmt.Sprintf(`## :mailbox_with_mail:質問が届きました

%v

質問日時：%v `,
		q, time.Now().Format("2006/01/02 15:04"))

	err = traq.PostWebhookOrPrint(message)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to send webhook: %v", err.Error()))
	}

	return c.JSON(http.StatusOK, PostQuestionResponse{
		Id:        question.Id,
		Question:  question.Question,
		CreatedAt: question.CreatedAt,
	})
}

func (h *questionHandler) GetQuestionByIdHandler(c echo.Context) error {
	id, err := uuid.Parse(c.Param("questionId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "bad request")
	}

	q, err := h.r.GetQuestionById(id)
	if errors.Is(err, repository.ErrNoRecord) {
		return echo.NewHTTPError(http.StatusNotFound, "no such question id")
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("failed to get question by id: %w", err))
	}
	return c.JSON(http.StatusOK, q)
}

func (h *questionHandler) GetQuestionsHandler(c echo.Context) error {
	var err error
	limit := 10
	offset := 0

	limitStr := c.QueryParam("limit")
	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil || limit <= 0 {
			return echo.NewHTTPError(http.StatusBadRequest, "bad limit")
		}
	}

	offsetStr := c.QueryParam("offset")
	if offsetStr != "" {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil || offset < 0 {
			return echo.NewHTTPError(http.StatusBadRequest, "bad offset")
		}
	}

	q, err := h.r.GetQuestions(limit, offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("failed to get questions: %w", err))
	}

	return c.JSON(http.StatusOK, q)
}

type PostAnswerRequest struct {
	Answer string `json:"answer"`
}

func (h *questionHandler) PostAnswerHandler(c echo.Context) error {
	id, err := uuid.Parse(c.Param("questionId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "bad id")
	}

	token, err := getToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to get token: %v", err))
	}

	userId, _, _, err := traq.GetMe(token)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to get me: %v", err))
	}

	var req PostAnswerRequest
	err = c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "bad request body")
	}
	q, err := h.r.CreateAnswer(id, req.Answer, userId)
	if errors.Is(err, repository.ErrNoRecord) {
		return echo.NewHTTPError(http.StatusNotFound, "no such id")
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("failed to post answer: %w", err))
	}

	return c.JSON(http.StatusOK, q)
}
