package domain

import "time"

type Orders struct {
	OrderID     int `gorm:"column:order_id;primaryKey;autoIncrement"`
	UserIDFK    int `gorm:"column:user_id_fk"`
	ProductIDFK int `gorm:"column:product_id_fk"`
	Quantity    int `gorm:"column:quantity"`
	Price       int `gorm:"column:price"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
