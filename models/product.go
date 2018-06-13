package models

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // need this to use db
)

// Product holds all the basic product info
type Product struct {
	gorm.Model
	Code  string
	Price float32
	Name  string
	Stock uint
}

// GetProduct returns the Product with given id.
func GetProduct(id uint, db *gorm.DB) *Product {
	p := &Product{}
	db.Table("products").Where("id = ?", id).First(p)
	return p
}

// GetProducts returns all Products.
func GetProducts(db *gorm.DB) ([]*Product, error) {
	var products []*Product
	db.Table("products").Find(&products)
	if len(products) == 0 {
		return products, errors.New("no products in db")
	}

	return products, nil
}
