package router

import (
	"net/http"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

type handler interface {
	HandleInsertBook(c *gin.Context)
}

func SetupRouter(r *gin.Engine, h handler) {
	r.Use(requestid.New())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	})

	r.POST("/book", h.HandleInsertBook)
}
