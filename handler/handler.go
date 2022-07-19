package handler

import (
	"context"
	"template-gin/model"

	"github.com/gin-gonic/gin"
)

type usecase interface {
	InsertBook(ctx context.Context, b model.Book) error
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
		SendBadRequestResponse(c, "invalid request payload", nil)
		return
	}

	err := h.uc.InsertBook(c.Request.Context(), req)
	if err != nil {
		SendInternalErrorResponse(c, "internal server error", nil)
		return
	}

	SendCreatedResponse(c, "book created", req)
}
