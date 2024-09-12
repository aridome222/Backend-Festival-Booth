package controller

import (
	"net/http"

	"github.com/aridome222/Backend-Festival-Booth/usecase"
	"github.com/gin-gonic/gin"
)

type CreateAccountController struct {
	uc usecase.CreateAccountUseCase
}

func NewCreateAccountController(uc usecase.CreateAccountUseCase) CreateAccountController {
	return CreateAccountController{
		uc: uc,
	}
}

func (con CreateAccountController) CreateAccount(ctx *gin.Context) {
	// JSONデータをInputDTOの形式に合わせて格納する変数
	var input usecase.CreateAccountInputDTO

	// リクエストから取り出したJSONデータを構造体にバインド
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// アカウント情報を保存するユースケースを実行
	createdAccount, err := con.uc.CreateAccount(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdAccount)
}
