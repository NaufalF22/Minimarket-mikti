package orderrepository

import "minimarket_mikti/model/domain"

type OrderRepositoryInterface interface {
	Create(order domain.Orders) (domain.Orders, error)
}
