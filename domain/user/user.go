package user

import (
	app "go-clean-architecture/util/application"

	"github.com/jinzhu/gorm"
)

// Model => user domain struct
type Model struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UsecaseInterface ...
type UsecaseInterface interface {
	List(ac app.CustomContextInterface, query app.Query) ([]Model, error)
	Create(ac app.CustomContextInterface, payload Model) (Model, error)
}

// RepoInterface ...
type RepoInterface interface {
	List(trx *gorm.DB, query app.Query) ([]Model, error)
	Create(trx *gorm.DB, payload Model) (Model, error)
}
