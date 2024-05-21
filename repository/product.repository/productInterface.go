package productrepository

import "minimarket_mikti/model/domain"

type ProductRepoInterface interface {
	GetProductID(id int) (domain.Products, error)
	GetAll() ([]domain.Products, error)
}
