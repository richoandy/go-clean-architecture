# Go-Clean-Architecture #
## Very simple demonstration on Clean Architecture for Go Lang: ##
- reduce the pointers usage to reduce complexity
- slim repository, fat usecase
- usecase should not interact with any 3rd party library dependencies
- "only test code that I write"


**stack**:
- Web Framework: *Echo*
- ORM: *GORM*
- Database: *Mysql*

**To do**:
1. add .env file to run the service
- `MYSQL_URI=root:@tcp(127.0.0.1:3306)/go-clean-architecture?charset=utf8mb4&parseTime=True&loc=Local`
- `PORT=1323`
2. create database in mysql named `go-clean-architecture`
3. create table named `users` with field id `integer`, email `string`, password `string`
4. command to run: `go run app/main.go`