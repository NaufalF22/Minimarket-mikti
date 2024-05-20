package service

import "minimarket_mikti/model/web"

type UserService interface {
	SaveUser(request web.UserServiceRequest) (map[string]interface{}, error)
}
