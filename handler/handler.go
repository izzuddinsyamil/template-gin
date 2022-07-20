package handler

import (
	"context"
	"log"
	"template-gin/model"

	"github.com/gin-gonic/gin"
)

type usecase interface {
	InsertBook(ctx context.Context, b model.Book) error
	GetBooks(ctx context.Context) (res []model.Book, err error)
}

type handler struct {
	uc usecase
}

func NewHandler(uc usecase) *handler {
	return &handler{
		uc: uc,
	}
}

func (h *handler) HandleInsertBook(c *gin.Context) {
	req := model.Book{}
	if err := c.ShouldBind(&req); err != nil {
		sendBadRequestResponse(c, "invalid request payload", nil)
		return
	}

	err := h.uc.InsertBook(c.Request.Context(), req)
	if err != nil {
		sendInternalErrorResponse(c, "internal server error", nil)
		return
	}

	sendCreatedResponse(c, "book created", req)
}

func (h *handler) HandleGetBooks(c *gin.Context) {
	books, err := h.uc.GetBooks(c.Request.Context())
	if err != nil {
		log.Printf("error while getting books: %v", err)
		sendInternalErrorResponse(c, "internal server error", nil)
		return
	}

	sendOkResponse(c, "ok", books)
}
