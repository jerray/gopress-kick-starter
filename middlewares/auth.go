package middlewares

import (
	"net/http"

	"github.com/fpay/gopress"
	starter "github.com/jerray/gopress-kick-starter"
)

const (
	globalToken = "Gopress"
)

// NewAuthMiddleware returns auth middleware.
func NewAuthMiddleware(userService starter.UserService) gopress.MiddlewareFunc {
	return func(next gopress.HandlerFunc) gopress.HandlerFunc {
		return func(c gopress.Context) error {
			token := c.Request().Header.Get("Token")
			if token != globalToken {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"code":    401,
					"message": "Unauthorized: invalid token",
				})
			}

			// user, err := userService.GetUserByID()

			WithUser(c, &starter.User{ID: 1})

			return next(c)
		}
	}
}

func WithUser(ctx gopress.Context, user *starter.User) {
	ctx.Set("user", user)
}

func ExtractUser(ctx gopress.Context) *starter.User {
	if u, ok := ctx.Get("user").(*starter.User); ok {
		return u
	}
	return nil
}
