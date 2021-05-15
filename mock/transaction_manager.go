package mock

import (
	"github.com/jinzhu/gorm"
)

// AcMock => mock struct
type TrxMock struct{}

// TrxStart => mock function
func (trxMock TrxMock) Begin() *gorm.DB { return &gorm.DB{} }

// TrxCommit => mock function
func (trxMock TrxMock) Commit(trx *gorm.DB) {}

// TrxRollback => mock function
func (trxMock TrxMock) Rollback(trx *gorm.DB) {}
