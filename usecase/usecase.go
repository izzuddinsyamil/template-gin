package usecase

type repo interface {
}

type usecase struct {
	repo repo
}

func NewUsecase(repo repo) *usecase {
	return &usecase{
		repo: repo,
	}
}
