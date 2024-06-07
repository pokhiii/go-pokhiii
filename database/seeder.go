package main

import (
	"encoding/json"
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Product represents a product in the eCommerce store
type Product struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Slug        string  `json:"slug" gorm:"unique"`
}

func main() {
	// Read seed data from JSON file
	seedData, err := os.ReadFile("database/seeders/products.json")
	if err != nil {
		fmt.Println("Error reading seed data:", err)
		return
	}

	// Unmarshal JSON data into slice of products
	var products []Product
	if err := json.Unmarshal(seedData, &products); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// Connect to SQLite database
	db, err := gorm.Open(sqlite.Open("products.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Seed the database with initial data
	for _, p := range products {
		if err := db.Create(&p).Error; err != nil {
			fmt.Println("Error seeding database:", err)
			return
		}
	}

	fmt.Println("Database seeded successfully")
}
