package router

import (
	"github.com/aridome222/Backend-Festival-Booth/api/controller"
	"github.com/aridome222/Backend-Festival-Booth/infrastructure/repository"
	"github.com/aridome222/Backend-Festival-Booth/usecase"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) {
	// get login
	loginRepository := repository.NewAccountRepository(db)
	loginUseCase := usecase.NewLoginUseCase(loginRepository)
	loginController := controller.NewLoginController(loginUseCase)

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

	// save profile
	profileRepository := repository.NewProfileRepository(db)
	saveProfileUseCase := usecase.NewSaveProfileUseCase(profileRepository)
	saveProfileController := controller.NewSaveProfileController(saveProfileUseCase)

	// get profiles
	getProfileListUseCase := usecase.NewGetProfileListUseCase(profileRepository)
	getProfileListController := controller.NewGetProfileListController(getProfileListUseCase)

	// get profile/:name
	getProfileUseCase := usecase.NewGetProfileUseCase(profileRepository)
	getProfileController := controller.NewGetProfileController(getProfileUseCase)

	// save comment
	commentRepository := repository.NewCommentRepository(db)
	saveCommentUseCase := usecase.NewSaveCommentUseCase(commentRepository, productRepository)
	saveCommentController := controller.NewSaveCommentController(saveCommentUseCase)

	// create account
	accountRepository := repository.NewAccountRepository(db)
	createAccountUseCase := usecase.NewCreateAccountUseCase(accountRepository)
	createAccountController := controller.NewCreateAccountController(createAccountUseCase)

	r := gin.Default()
	// TODO: cors設定を適宜調整
	r.Use(cors.New(cors.Config{
		// 許可するHTTPメソッド一覧
		AllowMethods: []string{
			"POST",
			"GET",
		},
		// 許可するHTTPリクエストヘッダ一覧
		// AllowHeaders: []string{
		//  "Access-Control-Allow-Credentials",
		// 	"Origin",
		// 	"Content-Type",
		// },
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: true,
	}))

	// /login
	r.POST("/login", loginController.Login)

	// /products
	r.POST("/products", saveProductController.SaveProduct)
	r.GET("/products", getProductListController.GetProductList)
	r.GET("/products/:user_name", getProductController.GetProduct)

	// /profiles
	r.POST("/profiles", saveProfileController.SaveProfile)
	r.GET("/profiles", getProfileListController.GetProfileList)
	r.GET("/profiles/:name", getProfileController.GetProfile)

	// /comments
	r.POST("/comments", saveCommentController.SaveComment)

	// /accounts
	r.POST("/accounts", createAccountController.CreateAccount)

	r.Run() // 0.0.0.0:8080 でサーバーを立てます。
}
