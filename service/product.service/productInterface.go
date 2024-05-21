package productservice

import (
	"minimarket_mikti/model/entity"
	"minimarket_mikti/model/web"
)

type ProductServiceInterface interface {
	GetProductID(id int) (entity.ProductEntity, error)
	GetAllProduct() ([]entity.ProductEntity, error)
	BuyProduct(id int, req web.OrderRequest) (map[string]interface{}, error)
}
