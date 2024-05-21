package productservice

import (
	"errors"
	"minimarket_mikti/model/domain"
	"minimarket_mikti/model/entity"
	"minimarket_mikti/model/web"
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

func (pS *ProductService) BuyProduct(id int, req web.OrderRequest) (map[string]interface{}, error) {
	product, errProduct := pS.repo.GetProductID(id)

	if errProduct != nil {
		return nil, errProduct
	}

	newQuantity := product.Quantity - req.Quantity

	if req.Quantity > product.Quantity {
		return nil, errors.New("quantity to much")
	}

	dataUpdate := domain.Products{
		ProductID:  product.ProductID,
		SellerIDFK: product.SellerIDFK,
		Name:       product.Name,
		Quantity:   newQuantity,
		Price:      product.Price,
	}

	_, errData := pS.repo.BuyProduct(dataUpdate)

	if errData != nil {
		return nil, errData
	}

	return nil, nil
}
