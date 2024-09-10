package router

import (
	"github.com/aridome222/Backend-Festival-Booth/api/controller"
	"github.com/aridome222/Backend-Festival-Booth/infrastructure/repository"
	"github.com/aridome222/Backend-Festival-Booth/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) {
	productRepository := repository.NewProductRepository(db)
	saveProductUseCase := usecase.NewSaveProductUseCase(productRepository)
	saveProductController := controller.NewSaveProductController(saveProductUseCase)

	r := gin.Default()
	// TODO: cors設定を追記

	r.POST("/products", saveProductController.SaveProduct)

	r.Run() // 0.0.0.0:8080 でサーバーを立てます。
}
