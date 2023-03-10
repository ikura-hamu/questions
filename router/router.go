package router

import (
	"github.com/ikura-hamu/questions/repository/impl"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/srinathgs/mysqlstore"
)

func SetUp(e *echo.Echo, db *sqlx.DB) {
	store, err := mysqlstore.NewMySQLStoreFromConnection(db.DB, "session", "/", 60*60*24*14, []byte("secret-token"))
	if err != nil {
		panic(err)
	}

	qh := NewQuestionHandler(impl.NewQuestionRepository(db))

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(session.Middleware(store))

	api := e.Group("/api")

	oauth2 := api.Group("/oauth2")
	oauth2.GET("/authorize", AuthorizeHandler)
	oauth2.GET("/callback", CallbackHandler)

	question := api.Group("/question")

	question.POST("", qh.PostQuestionHandler)
	question.GET("", qh.GetQuestionsHandler)
	question.GET("/:questionId", qh.GetQuestionByIdHandler)

	question.POST("/:questionId/answer", qh.PostAnswerHandler, CheckTraqLoginMiddleware)
}
