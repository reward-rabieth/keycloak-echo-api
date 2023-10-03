package handlers

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/reward-rabieth/Authclk/domain/entities"
	"github.com/reward-rabieth/Authclk/use_cases/productsuc"
	"net/http"
)

type CreateProductUseCase interface {
	CreateProduct(ctx context.Context, product productsuc.CreateProductRequest) (*productsuc.CreateProductResponse, error)
}

type GetProductUseCase interface {
	GetProducts(ctx context.Context) []entities.Product
}

func CreateProductHandler(useCase CreateProductUseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		var ctx = c.Request().Context()
		var request = productsuc.CreateProductRequest{}
		err := c.Bind(&request)
		if err != nil {
			return errors.Wrap(err, "unable to parse the incoming request")
		}
		response, err := useCase.CreateProduct(ctx, request)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusCreated, response)

	}
}
func GetProductsHandler(useCase GetProductUseCase) echo.HandlerFunc {

	return func(c echo.Context) error {
		var ctx = c.Request().Context()
		products := useCase.GetProducts(ctx)

		return c.JSON(http.StatusOK, products)
	}

}
