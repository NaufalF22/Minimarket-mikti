package productrepository

import (
	"minimarket_mikti/model/domain"

	"gorm.io/gorm"
)

type ProductRepo struct {
	DB *gorm.DB
}

func NewProductRepo(db *gorm.DB) *ProductRepo {
	return &ProductRepo{
		DB: db,
	}
}

func (pR *ProductRepo) GetProductID(id int) (domain.Products, error) {
	var product domain.Products

	if err := pR.DB.Where("product_id = ?", id).Take(&product).Error; err != nil {
		return domain.Products{}, err
	}

	return product, nil
}

func (pR *ProductRepo) GetAll() ([]domain.Products, error) {
	var products []domain.Products

	if err := pR.DB.Find(&products).Error; err != nil {
		return []domain.Products{}, err
	}

	return products, nil
}

func (pR *ProductRepo) BuyProduct(product domain.Products) (domain.Products, error) {
	if err := pR.DB.Model(domain.Products{}).Where("product_id = ?", product.ProductID).Updates(&product).Error; err != nil {
		return domain.Products{}, err
	}

	return product, nil
}
