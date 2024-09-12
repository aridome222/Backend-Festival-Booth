package usecase

import (
	"github.com/aridome222/Backend-Festival-Booth/domain"
)

type GetCommentUseCase struct {
	repo domain.CommentRepository
}

type GetCommentInputDTO struct {
	ProductID string
	Message   string
}

type GetCommentOutputDTO struct {
	ID        string `json:"id"`
	ProductID string `json:"product_id"`
	Message   string `json:"message"`
}

func NewGetCommentUseCase(repo domain.CommentRepository) GetCommentUseCase {
	return GetCommentUseCase{
		repo: repo,
	}
}

func (uc GetCommentUseCase) GetComment(input GetCommentInputDTO) ([]GetCommentOutputDTO, error) {
	var comments []domain.Comment
	var err error

	// ProductIDでコメントを取得
	comments, err = uc.repo.FindByProductID(input.ProductID)
	if err != nil {
		return nil, err
	}

	/* domain.Comment構造体のスライスcommentsを、GetCommentOUtputDTO構造体のスライスに変換 */
	outputSlice := []GetCommentOutputDTO{}
	// 各domain.CommentをGetCommentOutputDTOに変換
	for _, comment := range comments {
		outputSlice = append(outputSlice, GetCommentOutputDTO{
			ID:        comment.ID,
			ProductID: comment.ProductID,
			Message:   comment.Message,
		})
	}

	return outputSlice, nil
}
