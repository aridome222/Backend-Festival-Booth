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

	// get products
	getProductUseCase := usecase.NewGetProductUseCase(productRepository)
	getProductController := controller.NewGetProductController(getProductUseCase)

	// save profile
	profileRepository := repository.NewProfileRepository(db)
	saveProfileUseCase := usecase.NewSaveProfileUseCase(profileRepository)
	saveProfileController := controller.NewSaveProfileController(saveProfileUseCase)

	// get profiles
	getProfileUseCase := usecase.NewGetProfileUseCase(profileRepository)
	getProfileController := controller.NewGetProfileController(getProfileUseCase)

	r := gin.Default()
	// TODO: cors設定を追記

	r.POST("/products", saveProductController.SaveProduct)
	r.GET("/products", getProductController.GetProduct)
	r.POST("/profiles", saveProfileController.SaveProfile)
	r.GET("/profiles", getProfileController.GetProfile)

	r.Run() // 0.0.0.0:8080 でサーバーを立てます。
}
