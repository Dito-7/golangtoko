package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type OrderItem struct {
	ID              string `gorm:"size:36;not null;uniqueIndex;primary_key" json:"id"`
	Order           Order
	OrderID         string `gorm:"size:36;index" json:"order_id"`
	Product         Product
	ProductID       string `gorm:"size:36;index" json:"product_id"`
	Qty             int
	BasePrice       decimal.Decimal `gorm:"type:decimal(16,2)" json:"base_price"`
	BaseTotal       decimal.Decimal `gorm:"type:decimal(16,2)" json:"base_total"`
	TaxAmount       decimal.Decimal `gorm:"type:decimal(16,2)" json:"tax_amount"`
	TaxPercent      decimal.Decimal `gorm:"type:decimal(10,2)" json:"tax_percent"`
	DiscountAmount  decimal.Decimal `gorm:"type:decimal(16,2)" json:"discount_amount"`
	DiscountPercent decimal.Decimal `gorm:"type:decimal(10,2)" json:"discount_percent"`
	SubTotal        decimal.Decimal `gorm:"type:decimal(16,2)" json:"sub_total"`
	Sku             string          `gorm:"size:36;index" json:"sku"`
	Name            string          `gorm:"size:100" json:"name"`
	Weight          decimal.Decimal `gorm:"type:decimal(16,2)" json:"weight"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
