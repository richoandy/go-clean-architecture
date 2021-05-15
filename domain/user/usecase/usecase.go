package usecase

import (
	"go-clean-architecture/domain/user"
	"go-clean-architecture/util/application"
	TransactionManager "go-clean-architecture/util/transaction_manager"
)

type usecaseHandler struct {
	UserRepo   user.RepoInterface
	TrxManager TransactionManager.ITrxManager
}

// New => initialize User Usecase
func New(trxManager TransactionManager.ITrxManager, userRepo user.RepoInterface) user.UsecaseInterface {
	return usecaseHandler{
		UserRepo:   userRepo,
		TrxManager: trxManager,
	}
}

/**
Methods for User Usecase
usecases don't interact directly with libraries / 3rd party dependencies
*/

func (uc usecaseHandler) List(query application.Query) ([]user.User, error) {
	trx := uc.TrxManager.Begin()

	result, error := uc.UserRepo.List(trx, query)
	if error != nil {
		uc.TrxManager.Rollback(trx)
		return nil, error
	}

	uc.TrxManager.Commit(trx)
	return result, nil
}

func (uc usecaseHandler) Create(payload user.User) (user.User, error) {
	trx := uc.TrxManager.Begin()

	result, error := uc.UserRepo.Create(trx, payload)
	if error != nil {
		uc.TrxManager.Rollback(trx)
		return payload, error
	}

	uc.TrxManager.Commit(trx)
	return result, nil
}
