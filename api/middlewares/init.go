package middlewares

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/reward-rabieth/Authclk/infrastucture/identity"
	"github.com/reward-rabieth/Authclk/shared/enums"
	"log"
)

func InitEchoMiddleware(e *echo.Echo, initPublicRoutes func(e *echo.Echo), initProtectedRoutes func(e *echo.Echo)) {
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Get the request ID that was added by the RequestID middleware
			requestID := c.Response().Header().Get(echo.HeaderXRequestID)
			fmt.Println(requestID)
			// Create a new context and add the request ID to it
			ctx := context.WithValue(context.Background(), enums.ContextKeyRequestID, requestID)
			c.Set("user-context", ctx)
			return next(c)
		}
	})
	// Routes that don't require a JWT token
	initPublicRoutes(e)
	tokenRetroSpector := identity.NewIdentityManger()
	e.Use(CheckToken(tokenRetroSpector))
	//Routes that require authorization and authentication
	initProtectedRoutes(e)
	log.Println("echo middlewares initialized")

}
