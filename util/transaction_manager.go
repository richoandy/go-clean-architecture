package util

import (
	"github.com/jinzhu/gorm"
)

// CustomContext => struct usecase
type TrxManager struct {
	Db *gorm.DB
}

// CustomContextInterface ...
type ITrxManager interface {
	Begin() *gorm.DB
	Commit(trx *gorm.DB)
	Rollback(trx *gorm.DB)
}

// New => initialize application context
func New(db *gorm.DB) ITrxManager {
	return TrxManager{
		Db: db,
	}
}

// TrxStart => abstract transaction begin action for GORM ORM
func (trxManager TrxManager) Begin() *gorm.DB {
	trx := trxManager.Db.Begin()
	return trx
}

// TrxCommit => abstract transaction commit action for GORM ORM
func (trxManager TrxManager) Commit(trx *gorm.DB) {
	trx.Commit()
}

// TrxRollback => abstract transaction rollback action for GORM ORM
func (trxManager TrxManager) Rollback(trx *gorm.DB) {
	trx.Rollback()
}
