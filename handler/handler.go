package handler

import (
	"context"
	"log"
	"template-gin/model"

	"github.com/gin-gonic/gin"
)

type usecase interface {
	InsertBook(ctx context.Context, b model.Book) error
	GetBooks(ctx context.Context, title string) (res []model.Book, err error)
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
	req := insertBookRequest{}
	if err := c.ShouldBind(&req); err != nil {
		log.Printf("error when binding request %v\n", err)
		sendBadRequestResponse(c, "invalid request payload", nil)
		return
	}

	err := h.uc.InsertBook(c.Request.Context(), model.Book{
		Title:       req.Title,
		ReleaseYear: req.ReleaseYear,
	})
	if err != nil {
		sendInternalErrorResponse(c, "internal server error", nil)
		return
	}

	sendCreatedResponse(c, "book created", req)
}

func (h *handler) HandleGetBooks(c *gin.Context) {
	req := getBooksRequest{}
	if err := c.ShouldBind(&req); err != nil {
		sendBadRequestResponse(c, "invalid request payload", nil)
		return
	}

	books, err := h.uc.GetBooks(c.Request.Context(), req.Title)
	if err != nil {
		log.Printf("error while getting books: %v", err)
		sendInternalErrorResponse(c, "internal server error", nil)
		return
	}

	sendOkResponse(c, "ok", books)
}
