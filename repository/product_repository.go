package repository

import "minimarket_mikti/model/domain"

type ProductRepository interface {
	CreateProduct(product domain.Products) (domain.Products, error)
}
