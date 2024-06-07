package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Product represents a product in the eCommerce store
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Slug        string  `json:"slug"`
}

var products = []Product{
	{ID: 1, Name: "Hoodie with Pocket", Description: "Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Vestibulum tortor quam, feugiat vitae, ultricies eget, tempor sit amet, ante. Donec eu libero sit amet quam egestas semper. Aenean ultricies mi vitae est. Mauris placerat eleifend leo.", Price: 45.00, Slug: "hoodie-with-pocket"},
	{ID: 2, Name: "Hoodie with Zipper", Description: "Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Vestibulum tortor quam, feugiat vitae, ultricies eget, tempor sit amet, ante. Donec eu libero sit amet quam egestas semper. Aenean ultricies mi vitae est. Mauris placerat eleifend leo.", Price: 45.00, Slug: "hoodie-with-zipper"},
	{ID: 3, Name: "Long Sleeve Tee", Description: "Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Vestibulum tortor quam, feugiat vitae, ultricies eget, tempor sit amet, ante. Donec eu libero sit amet quam egestas semper. Aenean ultricies mi vitae est. Mauris placerat eleifend leo.", Price: 25.00, Slug: "long-sleeve-tee"},
}

// getProducts returns a list of products
func getProducts(c *gin.Context) {
	c.JSON(http.StatusOK, products)
}

// getProduct returns the details of a single product by slug
func getProduct(c *gin.Context) {
	slug := c.Param("product-slug")
	for _, product := range products {
		if product.Slug == slug {
			c.JSON(http.StatusOK, product)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
}

func main() {
	router := gin.Default()

	router.GET("/products", getProducts)
	router.GET("/product/:product-slug", getProduct)

	// Custom 404 handler
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
	})

	router.Run(":8080")
}
