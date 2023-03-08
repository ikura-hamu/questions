package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/ikura-hamu/questions/router"
)

func main() {
	db, err := sqlx.Open("mysql", getDSN())
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		log.Println(i)
		if err := db.DB.Ping(); err == nil {
			break
		} else if i == 9 {
			log.Fatal(err)
		}

		time.Sleep(time.Second * time.Duration(i+1))
	}

	e := echo.New()

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "pong!!!!!")
	})

	router.SetUp(e, db)

	e.Logger.Fatal(e.Start(":3000"))
}

func getDSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=True&loc=Asia%%2FTokyo&charset=utf8mb4",
		getEnvOrDefault("DB_USER", "root"),
		getEnvOrDefault("DB_PASS", "password"),
		getEnvOrDefault("DB_HOST", "db"),
		getEnvOrDefault("DB_PORT", "3306"),
		getEnvOrDefault("DB_NAME", "app"), // TODO: サービス名に変える
	)
}

func getEnvOrDefault(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return defaultValue
}
