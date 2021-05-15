package http

import (
	"go-clean-architecture/domain/user"
	"go-clean-architecture/util"
	app "go-clean-architecture/util/application"

	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type httpHandler struct {
	Usecase user.UsecaseInterface
	Db      *gorm.DB
}

// AddHandler => handler
func AddHandler(e *echo.Echo, db *gorm.DB, usecase user.UsecaseInterface) {
	userHandler := httpHandler{
		Usecase: usecase,
		Db:      db,
	}

	e.GET("/api/users", userHandler.list)
	e.POST("/api/users", userHandler.create)
}

func (h httpHandler) list(c echo.Context) error {
	trx := util.TrxManager{
		Db: h.Db,
	}

	query := c.QueryParam("query")
	queryCtx := app.Query{
		Query: query,
	}

	res, err := h.Usecase.List(trx, queryCtx)
	if err != nil {
		return app.JsonResponse(
			c,
			http.StatusInternalServerError,
			"internal server error",
			nil,
		)
	}

	return app.JsonResponse(c, http.StatusOK, "ok", res)
}

func (h httpHandler) create(c echo.Context) error {
	trx := util.TrxManager{
		Db: h.Db,
	}

	payload := user.User{}
	err := c.Bind(&payload)
	if err != nil {
		return app.JsonResponse(
			c,
			http.StatusUnprocessableEntity,
			"failed to bind payload",
			payload,
		)
	}

	res, err := h.Usecase.Create(trx, payload)
	if err != nil {
		return app.JsonResponse(
			c,
			http.StatusInternalServerError,
			"internal server error",
			payload,
		)
	}

	return app.JsonResponse(c, http.StatusOK, "ok", res)
}
