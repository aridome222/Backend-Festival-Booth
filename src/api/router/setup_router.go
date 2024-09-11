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
	getProductListUseCase := usecase.NewGetProductListUseCase(productRepository)
	getProductListController := controller.NewGetProductListController(getProductListUseCase)

	// get product/:user_name
	getProductUseCase := usecase.NewGetProductUseCase(productRepository)
	getProductController := controller.NewGetProductController(getProductUseCase)

	// get profiles
	profileRepository := repository.NewProfileRepository(db)
	getProfileListUseCase := usecase.NewGetProfileListUseCase(profileRepository)
	getProfileListController := controller.NewGetProfileListController(getProfileListUseCase)

	// get profile/:name
	getProfileUseCase := usecase.NewGetProfileUseCase(profileRepository)
	getProfileController := controller.NewGetProfileController(getProfileUseCase)

	// save comment
	commentRepository := repository.NewCommentRepository(db)
	saveCommentUseCase := usecase.NewSaveCommentUseCase(commentRepository)
	saveCommentController := controller.NewSaveCommentController(saveCommentUseCase)

	r := gin.Default()
	// TODO: cors設定を追記

	r.POST("/products", saveProductController.SaveProduct)
	r.GET("/products", getProductListController.GetProductList)
	r.GET("/products/:user_name", getProductController.GetProduct)
	r.GET("/profiles", getProfileListController.GetProfileList)
	r.GET("/profiles/:name", getProfileController.GetProfile)
	r.POST("/comments", saveCommentController.SaveComment)

	r.Run() // 0.0.0.0:8080 でサーバーを立てます。
}
