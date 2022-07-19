package handler

import (
	"context"
	"net/http"
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
	if err := c.BindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := h.uc.InsertBook(c.Request.Context(), req)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusAccepted, &req)
}
