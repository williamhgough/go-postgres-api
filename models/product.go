package models

import (
	"errors"
	"strconv"

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
func GetProduct(idString string, db *gorm.DB) (*Product, error) {
	id, err := strconv.ParseUint(idString, 0, 64)
	if err != nil {
		return nil, err
	}

	p := &Product{}
	db.Table("products").Where("id = ?", id).First(p)
	// If there is no name, product has not been found.
	if p.Name == "" {
		return nil, errors.New("no product found with specified id")
	}

	return p, nil
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
