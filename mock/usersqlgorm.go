package mock

import (
	"go-clean-architecture/domain/user"
	"go-clean-architecture/util/application"

	"github.com/jinzhu/gorm"
)

// UserRepoMock => mock struct
type UserRepoMock struct{}

// UserData for mocking
var UserData = user.Model{
	ID:       1,
	Email:    "test@email.com",
	Password: "testpassword",
}

// Create => mock function
func (m UserRepoMock) Create(trx *gorm.DB, payload user.Model) (user.Model, error) {
	return UserData, nil
}

// List => mock function
func (m UserRepoMock) List(trx *gorm.DB, query application.Query) ([]user.Model, error) {
	return []user.Model{UserData}, nil
}
