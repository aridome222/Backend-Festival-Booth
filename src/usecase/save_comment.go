package usecase

import (
	"github.com/aridome222/Backend-Festival-Booth/domain"
	"github.com/oklog/ulid/v2"
)

type SaveCommentUseCase struct {
	repo        domain.CommentRepository
	productRepo domain.ProductRepository
	accountRepo domain.AccountRepository
}

type SaveCommentUseCaseInputDTO struct {
	ProductID string `json:"product_id" binding:"required"`
	UserName  string
	Message   string `json:"message" binding:"required,min=1,max=100"`
}

type SaveCommentUseCaseOutputDTO struct {
	ID        string `json:"id"`
	UserName  string `json:"user_name"`
	ProductID string `json:"product_id"`
	Message   string `json:"message"`
}

func NewSaveCommentUseCase(
	repo domain.CommentRepository,
	productRepo domain.ProductRepository,
	accountRepo domain.AccountRepository,
) SaveCommentUseCase {
	return SaveCommentUseCase{
		repo:        repo,
		productRepo: productRepo,
		accountRepo: accountRepo,
	}
}

func (uc SaveCommentUseCase) SaveComment(input SaveCommentUseCaseInputDTO) (SaveCommentUseCaseOutputDTO, error) {
	_, err := uc.productRepo.FindByID(input.ProductID)
	// TODO: productが見つからなかった場合以外のエラーハンドリングを記述
	if err != nil {
		return SaveCommentUseCaseOutputDTO{}, err
	}

	_, err = uc.accountRepo.FindByName(input.UserName)
	// TODO: accountが見つからなかった場合以外のエラーハンドリングを記述
	if err != nil {
		return SaveCommentUseCaseOutputDTO{}, err
	}

	comment := domain.NewComment(ulid.Make().String(), input.UserName, input.ProductID, input.Message)

	comment, err = uc.repo.Save(comment)
	if err != nil {
		return SaveCommentUseCaseOutputDTO{}, err
	}

	return SaveCommentUseCaseOutputDTO{
		ID:        comment.ID,
		UserName:  comment.UserName,
		ProductID: comment.ProductID,
		Message:   comment.Message,
	}, nil
}
