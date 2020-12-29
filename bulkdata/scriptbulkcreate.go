package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	// User domain
	"go-clean-architecture/domain/user"
	UserRepoForSQLGorm "go-clean-architecture/domain/user/repository/gorm"
	UserUsecase "go-clean-architecture/domain/user/usecase"

	"go-clean-architecture/util/application"
	app "go-clean-architecture/util/application"
	"go-clean-architecture/util/gormloader"

	"github.com/joho/godotenv"
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

	// user domain
	userRepoForSQLGorm := UserRepoForSQLGorm.New(mysqlSession)
	userUsecase := UserUsecase.New(userRepoForSQLGorm)

	ac := app.CustomContext{
		Db: mysqlSession,
	}

	csvFile, err := os.Open("bulkdata/users.csv")
	if err != nil {
		log.Fatalln("Could not open csv file", err)
	}

	csvData, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	jobs := make(chan user.User, len(csvData))
	results := make(chan user.User, len(csvData))

	workerNumber := 10
	for i := 0; i < workerNumber; i++ {
		go createUserWorker(userUsecase, ac, jobs, results)
	}

	for i, data := range csvData {
		if i != 0 { // label for csv
			newEmail := data[0]
			newPassword := data[1]
			newUser := user.User{
				Email:    newEmail,
				Password: newPassword,
			}
			jobs <- newUser
		}

	}
	close(jobs)

	for result := range results {
		fmt.Print(result)
	}
}

func createUserWorker(usecase user.UsecaseInterface, ac application.CustomContextInterface, jobs <-chan user.User, results chan<- user.User) {
	for job := range jobs {
		result, _ := usecase.Create(ac, job)
		results <- result
	}
}
