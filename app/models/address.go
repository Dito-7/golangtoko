package models

import (
	"time"
)

type Address struct {
	ID        string `gorm:"size:36;not null;uniqueIndex;primary_key" json:"id"`
	User      User
	UserID    string `gorm:"size:36;index" json:"user_id"`
	Name      string `gorm:"size:100" json:"name"`
	IsPrimary bool
	CityID    string `gorm:"size:36" json:"city_id"`
	Addres1   string `gorm:"size:255" json:"addres1"`
	Addres2   string `gorm:"size:255" json:"addres2"`
	Phone     string `gorm:"size:100" json:"phone"`
	Email     string `gorm:"size:100" json:"email"`
	PostCode  string `gorm:"size:100" json:"post_code"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
