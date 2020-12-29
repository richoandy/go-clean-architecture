package gorm

import (
	"go-clean-architecture/domain/user"
	"go-clean-architecture/util/application"

	"github.com/jinzhu/gorm"
)

type repoHandler struct {
	Db *gorm.DB
}

// New => initialize User DAO with interface
func New(db *gorm.DB) user.RepoInterface {
	return repoHandler{
		Db: db,
	}
}

/**
keep methods in repository as slim as possible,
don't handle error or do conditional here => usecase's responsibility
*/

func (r repoHandler) List(trx *gorm.DB, query application.Query) ([]user.User, error) {
	payload := []user.User{}
	res := trx.Find(&payload)
	return payload, res.Error
}

func (r repoHandler) Create(trx *gorm.DB, payload user.User) (user.User, error) {
	res := trx.Create(&payload)
	return payload, res.Error
}
