package productsuc

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/reward-rabieth/Authclk/domain/entities"
	"time"
)

//todo try implementing  the echo custom validator

type CreateProductRequest struct {
	Name  string  `validate:"required,min=3,max=38" `
	Price float32 `validate:"required"`
}

type CreateProductResponse struct {
	Product *entities.Product
}
type createProductUseCase struct {
	dataStore ProductsDataStorer
}

func NewProductUseCase(ds ProductsDataStorer) *createProductUseCase {
	return &createProductUseCase{
		dataStore: ds,
	}
}
func (uc *createProductUseCase) CreateProduct(ctx context.Context, request CreateProductRequest) (*CreateProductResponse, error) {
	var validate = validator.New()
	err := validate.Struct(request)
	if err != nil {
		return nil, err
	}

	var product = &entities.Product{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		Name:      request.Name,
		Price:     request.Price,
	}
	err = uc.dataStore.Create(product)
	if err != nil {
		return nil, err
	}

	var response = &CreateProductResponse{Product: product}
	return response, nil
}
