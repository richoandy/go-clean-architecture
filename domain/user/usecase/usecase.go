package usecase

import (
	"go-clean-architecture/domain/user"
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

func (u usecaseHandler) List(ac application.CustomContextInterface, query application.Query) ([]user.Model, error) {
	trx := ac.TrxStart()

	result, error := u.UserRepo.List(trx, query)
	if error != nil {
		ac.TrxRollback(trx)
		return nil, error
	}

	ac.TrxCommit(trx)
	return result, nil
}

func (u usecaseHandler) Create(ac application.CustomContextInterface, payload user.Model) (user.Model, error) {
	trx := ac.TrxStart()

	result, error := u.UserRepo.Create(trx, payload)
	if error != nil {
		ac.TrxRollback(trx)
		return payload, error
	}

	ac.TrxCommit(trx)
	return result, nil
}
