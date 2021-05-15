package http

import (
	TransactionManager "go-clean-architecture/util/transaction_manager"

	// User
	UserHttp "go-clean-architecture/domain/user/delivery/http"
	UserRepo "go-clean-architecture/domain/user/repository"
	UserUsecase "go-clean-architecture/domain/user/usecase"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type HttpLoader struct {
	router *echo.Echo
	db     *gorm.DB
}

type IHttpLoader interface {
	Load()
}

func New(echo *echo.Echo, sqlSession *gorm.DB) IHttpLoader {
	return HttpLoader{
		router: echo,
		db:     sqlSession,
	}
}

func (httpLoader HttpLoader) Load() {
	trxManager := TransactionManager.New(httpLoader.db)

	// user domain
	userRepo := UserRepo.New(httpLoader.db)
	userUsecase := UserUsecase.New(trxManager, userRepo)

	UserHttp.AddHandler(httpLoader.router, userUsecase)
}
