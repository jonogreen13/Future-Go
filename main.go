package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/yourusername/yourapp/docs" // Import the generated Swagger docs (replace with your actual module path)
)

// @title Future YOU Application API
// @version 1.0.0
// @description Implements the functionality needed to run the Future YOU App Calculations
// @contact.email guy.fletcher.773@gmail.com
// @termsOfService https://site.futureyou.tech/terms

func main() {
	// Create the Gin router
	r := gin.Default()

	// Swagger route to access Swagger documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Register route handlers (modularized)
	registerRoutes(r)

	// Custom error handlers
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "The requested resource could not be found."})
	})

	// Start the server
	r.Run(":8080") // This will listen on port 8080
}

// This function will register the modular routes
func registerRoutes(r *gin.Engine) {
	// Root route
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the Future YOU Application API!",
		})
	})

	// Example of standalone routes
	r.GET("/bar", Bar)
	r.GET("/ind", Index)

	// Group related routes like in blueprints
	standard := r.Group("/standard")
	{
		standard.GET("/example", func(c *gin.Context) {
			c.String(200, "Standard API Example")
		})
	}

	premium := r.Group("/premium")
	{
		premium.GET("/example", func(c *gin.Context) {
			c.String(200, "Premium API Example")
		})
	}
}

// @Summary Returns the URL for the bar route
// @Description Example endpoint for the bar route
// @Success 200 {string} string "Success"
// @Router /bar [get]
func Bar(c *gin.Context) {
	c.String(200, "The URL for this page is /bar")
}

// @Summary Returns the URL for the index route
// @Description Example endpoint for the index route
// @Success 200 {string} string "Success"
// @Router /ind [get]
func Index(c *gin.Context) {
	c.String(200, "The URL for this page is /ind")
}
