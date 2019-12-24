package webT

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var AuthHandler = func(ctx echo.Context) bool {
	return false
}
var AuthMiddleware = func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if AuthHandler(c) {
			if err := AuthSuccess(c); err != nil {
				return err
			}
			return next(c)
		} else {
			return AuthErr(c)
		}
	}
}
var AuthErr = func(c echo.Context) error {
	return c.String(http.StatusUnauthorized, "Unauthorized!")
}
var AuthSuccess = func(c echo.Context) error {
	return nil
}
