package usecase

import (
	"context"
	"template-gin/model"
)

func (uc *usecase) InsertBook(ctx context.Context, b model.Book) error {
	return uc.repo.InsertBook(ctx, b)
}
