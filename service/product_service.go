package service

import "minimarket_mikti/model/web"

type ProductService interface {
	CreateProduct(request web.StoreProductRequest) (map[string]interface{}, error)
}
