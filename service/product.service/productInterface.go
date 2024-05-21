package productservice

import "minimarket_mikti/model/entity"

type ProductServiceInterface interface {
	GetProductID(id int) (entity.ProductEntity, error)
	GetAllProduct() ([]entity.ProductEntity, error)
}
