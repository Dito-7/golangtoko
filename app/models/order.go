package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Order struct {
	ID                  string `gorm:"size:36;not null;uniqueIndex;primary_key" json:"id"`
	UserID              string `gorm:"size:36;index" json:"user_id"`
	User                User
	OrderItems          []OrderItem
	OrderCustomer       *OrderCustomer
	Code                string `gorm:"size:100" json:"code"`
	Status              int    `json:"status"`
	OrderNote           time.Time
	PaymentDue          time.Time
	PaymentStatus       string          `gorm:"size:100" json:"payment_status"`
	PaymentToken        string          `gorm:"size:100" json:"payment_token"`
	BaseTotalPrice      decimal.Decimal `gorm:"type:decimal(16,2)" json:"base_total_price"`
	TaxAmount           decimal.Decimal `gorm:"type:decimal(16,2)" json:"tax_amount"`
	TaxPercent          decimal.Decimal `gorm:"type:decimal(10,2)" json:"tax_percent"`
	DiscountAmount      decimal.Decimal `gorm:"type:decimal(16,2)" json:"discount_amount"`
	DiscountPercent     decimal.Decimal `gorm:"type:decimal(10,2)" json:"discount_percent"`
	ShippingCost        decimal.Decimal `gorm:"type:decimal(16,2)" json:"shipping_cost"`
	GrandTotal          decimal.Decimal `gorm:"type:decimal(16,2)" json:"grand_total"`
	Note                string          `gorm:"type:text" json:"note"`
	ShippingCourier     string          `gorm:"size:100" json:"shipping_courier"`
	ShippingServiceName string          `gorm:"size:100" json:"shipping_service_name"`
	ApprovedBy          string          `gorm:"size:36" json:"approved_by"`
	ApprovedAt          time.Time
	CancelledBy         string `gorm:"size:36" json:"cancelled_by"`
	CancelledAt         time.Time
	CancelationNote     string `gorm:"size:255" json:"cancelation_note"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt
}
