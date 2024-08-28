package seeders

import (
	"fmt"

	"github.com/Dito-7/golangtoko/database/fakers"
	"gorm.io/gorm"
)

type Seeder struct {
	Seeder interface{}
}

func RegisterSeeders(db *gorm.DB) []Seeder {
	return []Seeder{
		{Seeder: fakers.UserFaker(db)},
		{Seeder: fakers.ProductFaker(db)},
	}
}

func RunSeeders(db *gorm.DB) error {
	for _, seeder := range RegisterSeeders(db) {
		err := db.Create(seeder.Seeder).Error
		if err != nil {
			return err
		}
	}

	fmt.Println("Seeders successfully run")

	return nil
}
