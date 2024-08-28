package models

import "time"

type ProductImage struct {
	ID         string `gorm:"size:36;not null;uniqueIndex;primary_key" json:"id"`
	Product    Product
	ProductID  string `gorm:"size:36;index" json:"product_id"`
	Path       string `gorm:"type:text" json:"path"`
	ExtraLarge string `gorm:"type:text" json:"extra_large"`
	Large      string `gorm:"type:text" json:"large"`
	Medium     string `gorm:"type:text" json:"medium"`
	Small      string `gorm:"type:text" json:"small"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
