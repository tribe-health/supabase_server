package auth

import "github.com/labstack/echo/v4"

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		auth := c.Request().Header.Get("Authorization")
		if auth == "" {
			return next(c)
		}
		return nil
	}
}
