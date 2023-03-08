package router

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ikura-hamu/questions/repository/impl"
)

func SetUp(e *echo.Echo, db *sqlx.DB) {
	//qh := NewQuestionHandler(impl.NewQuestionRepository(db))

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	api := e.Group("/api")
	question := api.Group("/question")
}
