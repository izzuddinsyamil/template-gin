package usecase

import (
	"context"
	"template-gin/model"
)

func (uc *usecase) GetBooks(ctx context.Context) (res []model.Book, err error) {
	return uc.repo.GetBooks(ctx)
}
