package usecase

import "github.com/aridome222/Backend-Festival-Booth/domain"

type GetCommentUseCase struct {
	repo domain.CommentRepository
}

type GetCommentInputDTO struct {
	ProductID string `json:"product_id" binding:"required"`
	Message   string `json:"message" binding:"required,min=1,max=100"`
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

func (uc GetCommentUseCase) GetComment(input GetCommentInputDTO) (GetCommentOutputDTO, error) {
	var comment domain.Comment
	var err error
	comment, err = uc.repo.FindByProductID(input.ProductID)

	if err != nil {
		return GetCommentOutputDTO{}, err
	}

	return GetCommentOutputDTO{
		comment.ID,
		comment.ProductID,
		comment.Message,
	}, nil
}
