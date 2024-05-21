package productservice

import (
	"minimarket_mikti/model/entity"
	productrepository "minimarket_mikti/repository/product.repository"
)

type ProductService struct {
	repo productrepository.ProductRepoInterface
}

func NewProductService(repo productrepository.ProductRepoInterface) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (pS *ProductService) GetProductID(id int) (entity.ProductEntity, error) {
	product, errProduct := pS.repo.GetProductID(id)

	if errProduct != nil {
		return entity.ProductEntity{}, errProduct
	}

	return entity.ToProduct(product), nil
}

func (pS *ProductService) GetAllProduct() ([]entity.ProductEntity, error) {
	products, errProducts := pS.repo.GetAll()

	if errProducts != nil {
		return []entity.ProductEntity{}, errProducts
	}

	return entity.ToProductList(products), nil
}
