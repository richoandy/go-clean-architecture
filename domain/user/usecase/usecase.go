package usecase

import (
	"go-clean-architecture/domain/user"
	"go-clean-architecture/util"
	"go-clean-architecture/util/application"
)

type usecaseHandler struct {
	UserRepo user.RepoInterface
}

// New => initialize User Usecase
func New(userRepo user.RepoInterface) user.UsecaseInterface {
	return usecaseHandler{
		UserRepo: userRepo,
	}
}

/**
Methods for User Usecase
usecases don't interact directly with libraries / 3rd party dependencies
*/

func (u usecaseHandler) List(trxMgr util.ITrxManager, query application.Query) ([]user.User, error) {
	trx := trxMgr.Begin()

	result, error := u.UserRepo.List(trx, query)
	if error != nil {
		trxMgr.Rollback(trx)
		return nil, error
	}

	trxMgr.Commit(trx)
	return result, nil
}

func (u usecaseHandler) Create(trxMgr util.ITrxManager, payload user.User) (user.User, error) {
	trx := trxMgr.Begin()

	result, error := u.UserRepo.Create(trx, payload)
	if error != nil {
		trxMgr.Rollback(trx)
		return payload, error
	}

	trxMgr.Commit(trx)
	return result, nil
}
