package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	// User domain
	"go-clean-architecture/domain/user"
	UserRepo "go-clean-architecture/domain/user/repository"
	UserUsecase "go-clean-architecture/domain/user/usecase"

	SqlConnector "go-clean-architecture/util/sql_connector"
	TransactionManager "go-clean-architecture/util/transaction_manager"

	"github.com/joho/godotenv"
)

func main() {
	// load.env
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	// initialize mysql connection
	mysqlSession := SqlConnector.OpenConnection(os.Getenv("MYSQL_URI"))
	trxManager := TransactionManager.New(mysqlSession)

	// user domain
	userRepoGorm := UserRepo.New(mysqlSession)
	userUsecase := UserUsecase.New(trxManager, userRepoGorm)

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
		go createUserWorker(userUsecase, trxManager, jobs, results)
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

func createUserWorker(usecase user.UsecaseInterface, trxManager TransactionManager.ITrxManager, jobs <-chan user.User, results chan<- user.User) {
	for job := range jobs {
		result, _ := usecase.Create(job)
		results <- result
	}
}
