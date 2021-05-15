package user

import (
	"go-clean-architecture/util"
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
	List(trx util.ITrxManager, query app.Query) ([]User, error)
	Create(trx util.ITrxManager, payload User) (User, error)
}

// RepoInterface ...
type RepoInterface interface {
	List(trx *gorm.DB, query app.Query) ([]User, error)
	Create(trx *gorm.DB, payload User) (User, error)
}
