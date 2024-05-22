package service

import (
	"minimarket_mikti/helper"
	"minimarket_mikti/model/domain"
	"minimarket_mikti/model/web"
	"minimarket_mikti/repository"
)

type ProductServiceImpl struct {
	repository repository.ProductRepository
}

func NewProductService(repository repository.ProductRepository) *ProductServiceImpl {
	return &ProductServiceImpl{
		repository: repository,
	}
}

func (service *ProductServiceImpl) CreateProduct(request web.StoreProductRequest) (map[string]interface{}, error) {
	product := domain.Products{
		SellerIDFK: request.SellerIDFK,
		Name:       request.Name,
		Quantity:   request.Quantity,
		Price:      request.Quantity,
	}

	saveProduct, errSaveProduct := service.repository.CreateProduct(product)

	if errSaveProduct != nil {
		return nil, errSaveProduct
	}

	return helper.ResponseToJson{
		"name":     saveProduct.Name,
		"quantity": saveProduct.Quantity,
		"price":    saveProduct.Price,
	}, nil
}
