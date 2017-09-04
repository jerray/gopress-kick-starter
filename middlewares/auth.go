package middlewares

import (
	"net/http"

	"github.com/fpay/gopress"
)

const (
	globalToken = "Gopress"
)

// NewAuthMiddleware returns auth middleware.
func NewAuthMiddleware() gopress.MiddlewareFunc {
	return func(next gopress.HandlerFunc) gopress.HandlerFunc {
		return func(c gopress.Context) error {
			// Uncomment this line if this middleware requires accessing to services.
			// services := gopress.AppFromContext(c).Services()

			token := c.Request().Header.Get("Token")
			if token != globalToken {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"code":    401,
					"message": "Unauthorized: invalid token",
				})
			}

			return next(c)
		}
	}
}
