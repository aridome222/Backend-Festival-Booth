package main

import (
	"github.com/aridome222/Backend-Festival-Booth/api/controller"
	"github.com/aridome222/Backend-Festival-Booth/infrastructure"
	"github.com/aridome222/Backend-Festival-Booth/infrastructure/repository"
	"github.com/aridome222/Backend-Festival-Booth/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	db := infrastructure.ConnectDB()

	productRepository := repository.NewProductRepository(db)
	saveProductUseCase := usecase.NewSaveProductUseCase(productRepository)
	saveProductController := controller.NewSaveProductController(saveProductUseCase)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/products", saveProductController.SaveProduct)

	r.Run() // 0.0.0.0:8080 でサーバーを立てます。
}
