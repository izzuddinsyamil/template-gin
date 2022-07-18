package server

import "github.com/gin-gonic/gin"

func NewServer() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.ErrorLogger())

	return r
}
