package repository

import (
	"gorm.io/gorm"
	"minimarket_mikti/model/domain"
)

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{
		db: db,
	}
}

func (repository *ProductRepositoryImpl) CreateProduct(product domain.Products) (domain.Products, error) {
	err := repository.db.Create(&product).Error

	if err != nil {
		return domain.Products{}, err
	}

	return product, nil
}
