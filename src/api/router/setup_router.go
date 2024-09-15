package router

import (
	"time"

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

	// save answer
	saveAnswerUseCase := usecase.NewSaveAnswerUseCase(accountRepository)
	saveAnswerController := controller.NewSaveAnswerController(saveAnswerUseCase)

	// save product
	productRepository := repository.NewProductRepository(db)
	saveProductUseCase := usecase.NewSaveProductUseCase(productRepository, accountRepository)
	saveProductController := controller.NewSaveProductController(saveProductUseCase)

	// get products
	getProductListUseCase := usecase.NewGetProductListUseCase(productRepository, accountRepository)
	getProductListController := controller.NewGetProductListController(getProductListUseCase)

	// get product/:user_name
	getProductUseCase := usecase.NewGetProductUseCase(productRepository)
	getProductController := controller.NewGetProductController(getProductUseCase)

	// save profile
	profileRepository := repository.NewProfileRepository(db)
	saveProfileUseCase := usecase.NewSaveProfileUseCase(profileRepository, accountRepository)
	saveProfileController := controller.NewSaveProfileController(saveProfileUseCase)

	// get profiles
	getProfileListUseCase := usecase.NewGetProfileListUseCase(profileRepository, accountRepository)
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
			"OPTIONS",
		},
		// 許可するHTTPリクエストヘッダ一覧
		AllowHeaders: []string{
			// "Access-Control-Allow-Credentials",
			// "Access-Control-Allow-Headers",
			"Content-Type",
			// "Content-Length",
			// "Accept-Encoding",
			// "Authorization",
			// "accessToken",
			"Set-Cookie",
			"Cookie",
		},
		AllowOrigins: []string{
			"http://localhost:5173",
			"https://frontend-festival-booth.vercel.app",
		},
		// cookieなどの情報を必要とするかどうか
		// AllowCredentials: true,
		MaxAge: 24 * time.Hour,
	}))
	// r.Use(cors.Default())
	// // セッションCookieの設定
	// store := cookie.NewStore([]byte("secret"))
	// store.Options(sessions.Options{
	//    Secure:   false,
	//    HttpOnly: false})

	// ルーティングのグループ化
	authRouter := r.Group("/", middleware.AuthJWT())
	// /login
	r.POST("/login", loginController.Login)
	authRouter.GET("/login", controller.GetAccount)

	// /accounts
	r.POST("/accounts", createAccountController.CreateAccount)

	// /answers
	authRouter.POST("/answers", saveAnswerController.SaveAnswer)

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

	// health check用
	r.GET("/", controller.Health)

	r.Run() // 0.0.0.0:8080 でサーバーを立てます。
}
