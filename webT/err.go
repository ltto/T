package webT

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var ErrorHandler = func(e error, c echo.Context) {
	err := c.JSON(http.StatusOK, e.Error())
	if err != nil {
		panic(err)
	}
}
