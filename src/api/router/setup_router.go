package router

import (
	"github.com/aridome222/Backend-Festival-Booth/api/controller"
	"github.com/aridome222/Backend-Festival-Booth/api/middleware"
	"github.com/aridome222/Backend-Festival-Booth/infrastructure/repository"
	"github.com/aridome222/Backend-Festival-Booth/usecase"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) {
	// create account
	accountRepository := repository.NewAccountRepository(db)
	createAccountUseCase := usecase.NewCreateAccountUseCase(accountRepository)
	createAccountController := controller.NewCreateAccountController(createAccountUseCase)

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
	saveCommentUseCase := usecase.NewSaveCommentUseCase(commentRepository, productRepository, accountRepository)
	saveCommentController := controller.NewSaveCommentController(saveCommentUseCase)

	// get comment
	getCommentUseCase := usecase.NewGetCommentUseCase(commentRepository)
	getCommentController := controller.NewGetCommentController(getCommentUseCase)

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
			"http://localhost:5173",
		},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: true,
	}))

	// ルーティングのグループ化
	authRouter := r.Group("/", middleware.AuthJWT())

	// /login
	r.POST("/login", loginController.Login)
	authRouter.GET("/login", controller.GetAccount)

	// /accounts
	r.POST("/accounts", createAccountController.CreateAccount)

	// /profiles
	authRouter.POST("/profiles", saveProfileController.SaveProfile)
	r.GET("/profiles", getProfileListController.GetProfileList)
	r.GET("/profiles/:name", getProfileController.GetProfile)

	// /products
	authRouter.POST("/products", saveProductController.SaveProduct)
	r.GET("/products", getProductListController.GetProductList)
	r.GET("/products/:user_name", getProductController.GetProduct)

	// /comments
	authRouter.POST("/comments", saveCommentController.SaveComment)
	r.GET("/comments/:product_id", getCommentController.GetComment)

	r.Run() // 0.0.0.0:8080 でサーバーを立てます。
}
