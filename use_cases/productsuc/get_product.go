package productsuc

import (
	"context"
	"github.com/reward-rabieth/Authclk/domain/entities"
)

type getProductsUseCase struct {
	datastore ProductsDataStorer
}

func NewGetProductsUseCase(ds ProductsDataStorer) *getProductsUseCase {

	return &getProductsUseCase{
		datastore: ds,
	}
}

func (uc *getProductsUseCase) GetProducts(ctx context.Context) []entities.Product {
	all := uc.datastore.GetAll()

	return all
}
