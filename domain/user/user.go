package user

import (
	app "go-clean-architecture/util/application"

	"github.com/jinzhu/gorm"
)

// User => user domain struct
type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UsecaseInterface ...
type UsecaseInterface interface {
	List(query app.Query) ([]User, error)
	Create(payload User) (User, error)
}

// RepoInterface ...
type RepoInterface interface {
	List(trx *gorm.DB, query app.Query) ([]User, error)
	Create(trx *gorm.DB, payload User) (User, error)
}
