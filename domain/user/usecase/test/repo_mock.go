package test

import (
	"go-clean-architecture/domain/user"
	"go-clean-architecture/util/application"

	"github.com/jinzhu/gorm"
)

// UserRepoMock => mock struct
type UserRepoMock struct{}

// UserData for mocking
var UserDataMock = user.User{
	ID:       1,
	Email:    "test@email.com",
	Password: "testpassword",
}

// Create => mock function
func (m UserRepoMock) Create(trx *gorm.DB, payload user.User) (user.User, error) {
	return UserDataMock, nil
}

// List => mock function
func (m UserRepoMock) List(trx *gorm.DB, query application.Query) ([]user.User, error) {
	return []user.User{UserDataMock}, nil
}
