package repository

import "minimarket_mikti/model/domain"

type UserRepository interface {
	SaveUser(user domain.Users) (domain.Users, error)
}
