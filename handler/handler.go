package handler

type usecase interface{}

type handler struct {
	uc usecase
}

func NewHandler(uc usecase) *handler {
	return &handler{
		uc: uc,
	}
}
