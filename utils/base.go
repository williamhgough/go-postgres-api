package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // need this to use db
	"github.com/subosito/gotenv"
	"github.com/williamhgough/pql-api/models"
)

var db *gorm.DB

func init() {
	err := gotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}

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
	GetDB().Create(&models.Product{
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