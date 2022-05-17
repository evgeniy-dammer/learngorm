package entities

import (
	"fmt"
	"time"
)

type Product struct {
	Id       int `gorm:"primary_key, AUTO_INCREMENT"`
	Name     string
	Price    float64
	Quantity int
	Status   bool
	Created  time.Time
}

func (product *Product) TableName() string {
	return "product"
}

func (product Product) ToString() string {
	return fmt.Sprintf(
		"\tId: %d\n\tName: %s\n\tPrice: %0.1f\n\tQuantity: %d\n\tStatus: %t\n\tCreated: %s",
		product.Id,
		product.Name,
		product.Price,
		product.Quantity,
		product.Status,
		product.Created.Format("02/01/2006"),
	)
}
