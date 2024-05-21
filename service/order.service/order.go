package orderservice

import (
	"errors"
	"minimarket_mikti/helper"
	"minimarket_mikti/model/domain"
	"minimarket_mikti/model/web"
	orderrepository "minimarket_mikti/repository/order.repository"
	productrepository "minimarket_mikti/repository/product.repository"
)

type OrderService struct {
	Repo        orderrepository.OrderRepositoryInterface
	productRepo productrepository.ProductRepoInterface
}

func NewOrderService(repo orderrepository.OrderRepositoryInterface, pR productrepository.ProductRepoInterface) *OrderService {
	return &OrderService{
		Repo:        repo,
		productRepo: pR,
	}
}

func (oS *OrderService) Create(userId int, id int, req web.OrderRequest) (map[string]interface{}, error) {
	product, errProduct := oS.productRepo.GetProductID(id)

	if errProduct != nil {
		return nil, errProduct
	}

	totalPrice := product.Price * req.Quantity

	if req.Price < totalPrice {
		return nil, errors.New("money not enough")
	}

	newOrder := domain.Orders{
		UserIDFK:    userId,
		ProductIDFK: id,
		Quantity:    req.Quantity,
		Price:       req.Price,
	}

	saveOrder, errOrder := oS.Repo.Create(newOrder)

	if errOrder != nil {
		return nil, errOrder
	}

	data := helper.CustomResponse{
		"order_id":    saveOrder.OrderID,
		"user_id":     saveOrder.UserIDFK,
		"product_id":  saveOrder.ProductIDFK,
		"quantity":    saveOrder.Quantity,
		"total_price": saveOrder.Price,
	}

	return data, nil
}
