package database

import (
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // need this to use db
	"github.com/williamhgough/go-postgres-api/models"
)

var db *gorm.DB

func init() {
	// Need to allow time for postgres to migrate and seed DB before running.
	time.Sleep(10 * time.Second)

	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("POSTGRES_HOST")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)

	conn, err := gorm.Open("postgres", dbURI)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.Debug().AutoMigrate(&models.Product{})
	Seed()
}

// GetDB provides access to db instance from
// other packages
func GetDB() *gorm.DB {
	return db
}

// Seed the database with two Initial Products
func Seed() {
	// discarding error as we only want to check
	// []*Products length
	products, _ := models.GetProducts(db)

	if len(products) == 0 {
		db.Create(&models.Product{
			Name:  "Socks",
			Price: 5,
			Stock: 98,
			Code:  "001",
		}).Create(&models.Product{
			Name:  "Shirt",
			Price: 49,
			Stock: 13,
			Code:  "002",
		})
	}
}
