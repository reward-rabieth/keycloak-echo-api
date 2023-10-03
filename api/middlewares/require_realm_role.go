package middlewares

import (
	"context"
	"fmt"
	jwt2 "github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/reward-rabieth/Authclk/shared/enums"
	"github.com/reward-rabieth/Authclk/shared/jwt"
	"net/http"
)

func NewRequiresRealmRole(role string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Retrieve the context from the Echo context
			retrievedContext := c.Get("context").(context.Context)
			// Retrieve the claims from the context
			claims := retrievedContext.Value(enums.ContextKeyClaims).(jwt2.MapClaims)

			fmt.Println("claims ...", claims)
			jwtHelper := jwt.NewJwtHelper(claims)
			if !jwtHelper.IsUserInRealmRole(role) {
				return c.String(http.StatusUnauthorized, "role authorization failed")
			}
			return next(c)
		}
	}
}
