package mock

import (
	"github.com/jinzhu/gorm"
)

// AcMock => mock struct
type AcMock struct{}

// TrxStart => mock function
func (m AcMock) TrxStart() *gorm.DB { return &gorm.DB{} }

// TrxCommit => mock function
func (m AcMock) TrxCommit(trx *gorm.DB) {}

// TrxRollback => mock function
func (m AcMock) TrxRollback(trx *gorm.DB) {}
