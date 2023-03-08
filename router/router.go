package router

import (
	"github.com/ikura-hamu/questions/repository/impl"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetUp(e *echo.Echo, db *sqlx.DB) {
	qh := NewQuestionHandler(impl.NewQuestionRepository(db))

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	api := e.Group("/api")
	question := api.Group("/question")

	question.POST("", qh.PostQuestionHandler)
	question.GET("", qh.GetQuestionsHandler)
	question.GET("/:questionId", qh.GetQuestionByIdHandler)
	question.POST("/:questionId/answer", qh.PostAnswerHandler)
}
