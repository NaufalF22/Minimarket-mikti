package orderservice

import "minimarket_mikti/model/web"

type OrderServiceInterface interface {
	Create(userId int, id int, req web.OrderRequest) (map[string]interface{}, error)
}
