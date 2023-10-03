package dataStores

import (
	"github.com/google/uuid"
	"github.com/reward-rabieth/Authclk/domain/entities"
	"sort"
	"sync"
)

type ProductDataStore struct {
	Products map[uuid.UUID]entities.Product
	sync.Mutex
}

func NewProductStore() *ProductDataStore {
	return &ProductDataStore{
		Products: make(map[uuid.UUID]entities.Product),
	}
}

func (ds *ProductDataStore) Create(product *entities.Product) error {
	ds.Lock()
	ds.Products[product.ID] = *product
	ds.Unlock()
	return nil
}

func (ds *ProductDataStore) GetAll() []entities.Product {
	all := make([]entities.Product, 0, len(ds.Products))
	for _, value := range ds.Products {
		all = append(all, value)
	}
	//sorting them based on creeation
	sort.Slice(all, func(i, j int) bool {
		return all[i].CreatedAt.Before(all[j].CreatedAt)
	})

	return all
}
