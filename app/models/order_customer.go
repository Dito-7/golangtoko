package models

import "time"

type OrderCustomer struct {
	ID         string `gorm:"size:36;not null;uniqueIndex;primary_key" json:"id"`
	User       User
	UserID     string `gorm:"size:36;index" json:"user_id"`
	Order      Order
	OrderID    string `gorm:"size:36;index" json:"order_id"`
	FirstName  string `gorm:"size:100;not null" json:"first_name"`
	LastName   string `gorm:"size:100;not null" json:"last_name"`
	CityID     string `gorm:"size:36" json:"city_id"`
	ProvinceID string `gorm:"size:36" json:"province_id"`
	Address1   string `gorm:"size:100" json:"address1"`
	Address2   string `gorm:"size:100" json:"address2"`
	Phone      string `gorm:"size:50" json:"phone"`
	Email      string `gorm:"size:100" json:"email"`
	PostCode   string `gorm:"size:100" json:"post_code"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
