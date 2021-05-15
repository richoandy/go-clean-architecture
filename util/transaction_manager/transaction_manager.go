package transaction_manager

import (
	"github.com/jinzhu/gorm"
)

// CustomContext => struct usecase
type trxManager struct {
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
	return trxManager{
		Db: db,
	}
}

// TrxStart => abstract transaction begin action for GORM ORM
func (trxMgr trxManager) Begin() *gorm.DB {
	trx := trxMgr.Db.Begin()
	return trx
}

// TrxCommit => abstract transaction commit action for GORM ORM
func (trxMgr trxManager) Commit(trx *gorm.DB) {
	trx.Commit()
}

// TrxRollback => abstract transaction rollback action for GORM ORM
func (trxMgr trxManager) Rollback(trx *gorm.DB) {
	trx.Rollback()
}
