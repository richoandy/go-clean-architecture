package main

import (
	"go-clean-architecture/app/http"
	SqlConnector "go-clean-architecture/util/sql_connector"

	"fmt"
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

	// initialize mysql connection
	sqlSession := SqlConnector.OpenConnection(os.Getenv("MYSQL_URI"))

	// initialize echo router
	e := echo.New()
	e.Use(middleware.CORS())

	// load HTTP router
	httpLoader := http.New(e, sqlSession)
	httpLoader.Load()

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
