package application

import "github.com/labstack/echo"

// Response => standarized response
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// JSONResponse => wrapper for JSON repsonse
func JsonResponse(c echo.Context, status int, message string, data interface{}) error {
	return c.JSON(status, Response{
		Status:  status,
		Message: message,
		Data:    data,
	})
}
