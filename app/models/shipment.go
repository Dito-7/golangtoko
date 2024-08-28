package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Shipment struct {
	ID          string `gorm:"size:36;not null;uniqueIndex;primary_key" json:"id"`
	User        User
	UserID      string `gorm:"size:36;index" json:"user_id"`
	Order       Order
	OrderID     string `gorm:"size:36;index" json:"order_id"`
	TrackNumber string `gorm:"size:255;index" json:"track_number"`
	Status      string `gorm:"size:36;index" json:"status"`
	TotalQty    int
	TotalWeight decimal.Decimal `gorm:"type:decimal(10,2)" json:"total_weight"`
	FirstName   string          `gorm:"size:100;not null" json:"first_name"`
	LastName    string          `gorm:"size:100;not null" json:"last_name"`
	CityID      string          `gorm:"size:100" json:"city_id"`
	ProvinceID  string          `gorm:"size:100" json:"province_id"`
	Address1    string          `gorm:"size:100" json:"address1"`
	Address2    string          `gorm:"size:100" json:"address2"`
	Phone       string          `gorm:"size:100" json:"phone"`
	Email       string          `gorm:"size:100" json:"email"`
	PostCode    string          `gorm:"size:100" json:"post_code"`
	ShippedBy   string          `gorm:"size:36" json:"shipped_by"`
	ShippedAt   time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
