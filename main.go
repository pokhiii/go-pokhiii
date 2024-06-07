package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

var DB *gorm.DB
var err error

// Initialize the database
func initDB() {
	DB, err = gorm.Open(sqlite.Open("products.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	DB.AutoMigrate(&Product{})
}

// getProducts returns a list of products
func getProducts(c *gin.Context) {
	var products []Product
	DB.Find(&products)
	c.JSON(http.StatusOK, products)
}

// getProduct returns the details of a single product by slug
func getProduct(c *gin.Context) {
	slug := c.Param("product-slug")
	var product Product
	if result := DB.Where("slug = ?", slug).First(&product); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func main() {
	router := gin.Default()

	// Initialize the database
	initDB()

	router.GET("/products", getProducts)
	router.GET("/product/:product-slug", getProduct)

	// Custom 404 handler
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Page not found"})
	})

	router.Run(":8080")
}
