package main

import (
	"log"

	"github.com/Qy-wy/beauty_product.git/internal/handlers"
	"github.com/Qy-wy/beauty_product.git/internal/models"
	"github.com/Qy-wy/beauty_product.git/internal/repositories"
	"github.com/Qy-wy/beauty_product.git/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=user1 password=password dbname=beauty_products_2 port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection error:", err)
	}

	db.AutoMigrate(&models.Product{}, &models.Brand{})

	productRepo := repositories.NewProductRepository(db)
	brandRepo := repositories.NewBrandRepository(db)

	productService := services.NewProductService(productRepo)
	brandService := services.NewBrandService(brandRepo)

	productHandler := handlers.NewProductHandler(productService)
	brandHandler := handlers.NewBrandHandler(brandService)

	router := gin.Default()

	router.GET("/products", productHandler.GetAll)
	router.GET("/products/:id", productHandler.GetByID)
	router.POST("/products", productHandler.Create)
	router.PUT("/products/:id", productHandler.Update)
	router.DELETE("/products/:id", productHandler.Delete)

	router.GET("/brands", brandHandler.GetAll)
	router.GET("/brands/:id", brandHandler.GetByID)
	router.POST("/brands", brandHandler.Create)
	router.PUT("/brands/:id", brandHandler.Update)
	router.DELETE("/brands/:id", brandHandler.Delete)

	log.Println("Server running on port 8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
