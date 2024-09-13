package controller

import (
	"net/http"

	"github.com/aridome222/Backend-Festival-Booth/usecase"
	"github.com/gin-gonic/gin"
)

type SaveProfileController struct {
	uc usecase.SaveProfileUsecase
}

func NewSaveProfileController(uc usecase.SaveProfileUsecase) SaveProfileController {
	return SaveProfileController{
		uc: uc,
	}
}

func (con SaveProfileController) SaveProfile(ctx *gin.Context) {
	// JSONデータをInputDTOの形式に合わせて格納する変数
	var input usecase.SaveProfileInputDTO

	// リクエストから取り出したJSONデータを構造体にバインド
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// cookieからユーザー名を取得
	user_name, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	input.Name = user_name.(string)

	// 自己紹介を保存するメソッドを実行
	savedProfile, err := con.uc.SaveProfile(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, savedProfile)
}
