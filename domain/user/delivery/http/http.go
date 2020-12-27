package http

import (
	"go-clean-architecture/domain/user"
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
	ac := app.CustomContext{
		Db: h.Db,
	}

	from := c.QueryParam("from")
	to := c.QueryParam("to")
	query := app.Query{
		From: from,
		To:   to,
	}

	res, err := h.Usecase.List(ac, query)
	if err != nil {
		return app.JSONResponse(c, http.StatusInternalServerError, "internal server error", nil)
	}

	return app.JSONResponse(c, http.StatusOK, "ok", res)
}

func (h httpHandler) create(c echo.Context) error {
	ac := app.CustomContext{
		Db: h.Db,
	}

	payload := user.Model{}
	err := c.Bind(&payload)
	if err != nil {
		return app.JSONResponse(c, http.StatusUnprocessableEntity, "failed to bind payload", payload)
	}

	res, err := h.Usecase.Create(ac, payload)
	if err != nil {
		return app.JSONResponse(c, http.StatusInternalServerError, "internal server error", payload)
	}

	return app.JSONResponse(c, http.StatusOK, "ok", res)
}
