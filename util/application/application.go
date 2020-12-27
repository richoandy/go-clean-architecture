package application

import (
	"github.com/jinzhu/gorm"
)

// CustomContext => struct usecase
type CustomContext struct {
	Db *gorm.DB
}

// CustomContextInterface ...
type CustomContextInterface interface {
	TrxStart() *gorm.DB
	TrxCommit(trx *gorm.DB)
	TrxRollback(trx *gorm.DB)
}

// New => initialize application context
func New(db *gorm.DB) CustomContextInterface {
	return CustomContext{
		Db: db,
	}
}

// TrxStart => abstract transaction begin action for GORM ORM
func (cc CustomContext) TrxStart() *gorm.DB {
	trx := cc.Db.Begin()
	return trx
}

// TrxCommit => abstract transaction commit action for GORM ORM
func (cc CustomContext) TrxCommit(trx *gorm.DB) {
	trx.Commit()
}

// TrxRollback => abstract transaction rollback action for GORM ORM
func (cc CustomContext) TrxRollback(trx *gorm.DB) {
	trx.Rollback()
}
