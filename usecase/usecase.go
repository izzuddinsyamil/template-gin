package usecase

import (
	"context"
	"template-gin/model"
)

type repo interface {
	InsertBook(ctx context.Context, b model.Book) error
}

type usecase struct {
	repo repo
}

func NewUsecase(repo repo) *usecase {
	return &usecase{
		repo: repo,
	}
}
