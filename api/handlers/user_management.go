package handlers

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/reward-rabieth/Authclk/use_cases/usermgmtuc"
	"net/http"
)

type RegisterUseCase interface {
	Register(context.Context, usermgmtuc.RegisterRequest) (*usermgmtuc.RegisterResponse, error)
}

func RegisterHandler(useCase RegisterUseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		var ctx = c.Request().Context()
		var request = usermgmtuc.RegisterRequest{}
		err := c.Bind(&request)
		if err != nil {
			return errors.Wrap(err, "unable to bind the incoming request")
		}
		response, err := useCase.Register(ctx, request)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, response)

	}
}
