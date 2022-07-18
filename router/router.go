package router

import (
	"net/http"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.Use(requestid.New())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	})
}
