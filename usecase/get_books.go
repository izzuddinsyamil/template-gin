package usecase

import (
	"context"
	"template-gin/model"
)

func (uc *usecase) GetBooks(ctx context.Context, title string, startYear, endYear int) (res []model.Book, err error) {
	return uc.repo.GetBooks(ctx, title, startYear, endYear)
}
