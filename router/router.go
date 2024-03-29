package router

import (
	"net/http"
	"os"

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

	clientUrl := getEnvOrDefault("CLIENT_URL", "http://localhost:5173")

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{clientUrl},
		AllowMethods:     []string{http.MethodGet, http.MethodPost},
		AllowCredentials: true,
	}))
	e.Use(session.Middleware(store))

	api := e.Group("/api")

	api.GET("/me", GetMe, CheckTraqLoginMiddleware)

	oauth2 := api.Group("/oauth2")
	oauth2.GET("/authorize", AuthorizeHandler)
	oauth2.GET("/callback", CallbackHandler)

	question := api.Group("/question")

	question.POST("", qh.PostQuestionHandler)
	question.GET("", qh.GetAnsweredQuestionsHandler)
	question.GET("/:questionId", qh.GetQuestionByIdHandler)

	admin := api.Group("/admin")
	admin.Use(CheckTraqLoginMiddleware)
	admin.GET("/question", qh.GetQuestionsHandler)
	admin.POST("/question/:questionId/answer", qh.PostAnswerHandler)
}

func getEnvOrDefault(envKey string, defaultValue string) string {
	value, ok := os.LookupEnv(envKey)
	if !ok {
		return defaultValue
	}
	return value
}
