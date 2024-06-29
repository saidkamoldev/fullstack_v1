// package main

// import (
// 	"net/http"
// 	"github.com/gin-gonic/gin"
// 	"github.com/gin-contrib/cors"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// type Product struct {
// 	gorm.Model
// 	Name  string
// 	Age   string
// }

// func main() {
// 	router := gin.Default()

// 	// CORS sozlamalari
// 	config := cors.DefaultConfig()
// 	config.AllowOrigins = []string{"http://127.0.0.1:5500"}
// 	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"}
// 	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type"}
// 	router.Use(cors.New(config))

// 	// Database connection
// 	dsn := "host=localhost user=saidkamol password=123456 dbname=mygo port=5432 sslmode=disable"
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect database")
// 	}
// 	db.AutoMigrate(&Product{})

// 	// Router handlers
// 	router.GET("/products", func(c *gin.Context) {
// 		var products []Product
// 		db.Find(&products)
// 		c.JSON(http.StatusOK, products)
// 	})

// 	router.POST("/products", func(c *gin.Context) {
// 		var product Product
// 		if err := c.ShouldBindJSON(&product); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		db.Create(&product)
// 		c.JSON(http.StatusOK, product)
// 	})
   
	
// 	// Run the server
// 	router.Run(":8080")
// }

package main

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name  string
	Age   string
}

func main() {
	router := gin.Default()

	// CORS sozlamalari
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://127.0.0.1:5500"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type"}
	router.Use(cors.New(config))

	// Database connection
	dsn := "host=localhost user=saidkamol password=123456 dbname=mygo port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Product{})

	// Router handlers
	router.GET("/products", func(c *gin.Context) {
		var products []Product
		db.Find(&products)
		c.JSON(http.StatusOK, products)
	})

	router.POST("/products", func(c *gin.Context) {
		var product Product
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&product)
		c.JSON(http.StatusOK, product)
	})

	// Clear products function
	clearProducts := func() error {
		// Log the action for debugging
		log.Println("Clearing all products from the database...")

		// Debugging: Check the number of products before deletion
		var countBefore int64
		if err := db.Model(&Product{}).Count(&countBefore).Error; err != nil {
			log.Printf("Error counting products before deletion: %v\n", err)
			return err
		}
		log.Printf("Number of products before deletion: %d\n", countBefore)

		// Delete products
		if err := db.Where("1 = 1").Delete(&Product{}).Error; err != nil {
			log.Printf("Error deleting products: %v\n", err)
			return err
		}

		// Debugging: Check the number of products after deletion
		var countAfter int64
		if err := db.Model(&Product{}).Count(&countAfter).Error; err != nil {
			log.Printf("Error counting products after deletion: %v\n", err)
			return err
		}
		log.Printf("Number of products after deletion: %d\n", countAfter)

		return nil
	}

	// Clear products endpoint
	router.DELETE("/clear-products", func(c *gin.Context) {
		if err := clearProducts(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "All products deleted"})
	})

	// Run the server
	router.Run(":8080")
}
