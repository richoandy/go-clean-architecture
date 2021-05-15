package http

import (
	// User
	UserHttp "go-clean-architecture/domain/user/delivery/http"
	UserRepo "go-clean-architecture/domain/user/repository"
	UserUsecase "go-clean-architecture/domain/user/usecase"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type HttpLoader struct {
	db     *gorm.DB
	router *echo.Echo
}

type IHttpLoader interface {
	Load()
}

func New(sqlSession *gorm.DB, echo *echo.Echo) IHttpLoader {
	return HttpLoader{
		db:     sqlSession,
		router: echo,
	}
}

func (httpLoader HttpLoader) Load() {
	// user domain
	userRepo := UserRepo.New(httpLoader.db)
	userUsecase := UserUsecase.New(userRepo)
	UserHttp.AddHandler(httpLoader.router, httpLoader.db, userUsecase)
}
