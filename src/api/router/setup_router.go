package router

import (
	"github.com/aridome222/Backend-Festival-Booth/api/controller"
	"github.com/aridome222/Backend-Festival-Booth/infrastructure/repository"
	"github.com/aridome222/Backend-Festival-Booth/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) {
	// save product
	productRepository := repository.NewProductRepository(db)
	saveProductUseCase := usecase.NewSaveProductUseCase(productRepository)
	saveProductController := controller.NewSaveProductController(saveProductUseCase)

	// get users
	userRepository := repository.NewUserRepository(db)
	getUserUseCase := usecase.NewGetUserUseCase(userRepository)
	getUserController := controller.NewGetUserController(getUserUseCase)

	r := gin.Default()
	// TODO: cors設定を追記

	r.POST("/products", saveProductController.SaveProduct)
	r.GET("/users", getUserController.GetUser)

	r.Run() // 0.0.0.0:8080 でサーバーを立てます。
}
