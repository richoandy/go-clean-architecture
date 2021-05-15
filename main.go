package main

import (
	"go-clean-architecture/app/http"
	"go-clean-architecture/util/sql_connector"

	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// load.env
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	// initialize mysql connection with GORM ORM
	sqlSession, err := sql_connector.OpenSqlConnection(os.Getenv("MYSQL_URI"))
	if err != nil {
		log.Print(err)
		panic("failed to initialize connection to mysql database")
	}

	// initialize echo router
	e := echo.New()
	e.Use(middleware.CORS())

	// load HTTP router
	httpLoader := http.New(sqlSession, e)
	httpLoader.Load()

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
