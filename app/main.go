package main

import (
	"fmt"
	"go-clean-architecture/util/gormloader"

	// User domain
	UserRepoForSQLGorm "go-clean-architecture/domain/user/repository/gorm"
	UserUsecase "go-clean-architecture/domain/user/usecase"

	// Delivery over HTTP
	UserHttp "go-clean-architecture/domain/user/delivery/http"

	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func main() {
	// load.env
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	// initialize mysql connection with GORM ORM
	mysqlSession, err := gormloader.OpenConnection(os.Getenv("MYSQL_URI"))
	if err != nil {
		log.Print(err)
		panic("failed to initialize connection to mysql database")
	}

	// initialize echo router
	e := echo.New()

	// user domain
	userRepoForSQLGorm := UserRepoForSQLGorm.New(mysqlSession)
	userUsecase := UserUsecase.New(userRepoForSQLGorm)

	// HTTP delivery
	UserHttp.AddHandler(e, mysqlSession, userUsecase)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
