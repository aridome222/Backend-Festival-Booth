package controller

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/aridome222/Backend-Festival-Booth/usecase"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type LoginController struct {
	uc usecase.LoginUseCase
}

func NewLoginController(uc usecase.LoginUseCase) LoginController {
	return LoginController{
		uc: uc,
	}
}

func (con LoginController) Login(ctx *gin.Context) {
	var input usecase.LoginInputDTO

	// jsonデータを構造体にデコードしバインド
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ログイン処理に成功したか判定
	accountName, isLogin := con.uc.Login(input)
	if !isLogin {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": errors.New("failed to login").Error()})
		return
	}

	// JWTの発行
	claims := jwt.MapClaims{
		"user_name": accountName,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := os.Getenv("SECRET_KEY")
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": errors.New("failed to generate token").Error()},
		)
		return
	}

	cookie := new(http.Cookie)
	cookie.Value = tokenString

	ctx.SetSameSite((http.SameSiteNoneMode))

	// cookieにjwtをセット
	if product := os.Getenv(("PRODUCTION")); product == "" {
		ctx.SetCookie("jwt", tokenString, 3600, "/", "localhost", false, true)
	} else {
		ctx.SetCookie("jwt", tokenString, 3600, "/", "backend-festival-booth.onrender.com", true, true)
		ctx.SetCookie("token", tokenString, 3600, "/", "frontend-festival-booth.vercel.app/", true, true)
		ctx.SetCookie("jwt2", tokenString, 3600, "/", "localhost", true, true)
		ctx.SetCookie("jwtHttp", cookie.Value, 3600, "/", "backend-festival-booth.onrender.com", true, true)
	}
	ctx.SetCookie("test", "ok", 3000, "/", "backend-festival-booth.onrender.com", true, true)
	ctx.JSON(http.StatusOK, gin.H{
		"user_name": accountName,
	})
}
