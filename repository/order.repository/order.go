package orderrepository

import (
	"minimarket_mikti/model/domain"

	"gorm.io/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return OrderRepository{
		DB: db,
	}
}

func (oR *OrderRepository) Create(order domain.Orders) (domain.Orders, error) {
	if err := oR.DB.Create(&order).Error; err != nil {
		return domain.Orders{}, err
	}

	return order, nil
}
