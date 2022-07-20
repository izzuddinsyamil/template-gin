package router

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

type handler interface {
	HandleInsertBook(c *gin.Context)
	HandleGetBooks(c *gin.Context)
}

func SetupRouter(r *gin.Engine, h handler) {
	r.Use(requestid.New())

	r.POST("/book", h.HandleInsertBook)
	r.GET("/books", h.HandleGetBooks)
}
