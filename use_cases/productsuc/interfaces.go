package productsuc

import "github.com/reward-rabieth/Authclk/domain/entities"

type ProductsDataStorer interface {
	Create(products *entities.Product) error
	GetAll() []entities.Product
}
