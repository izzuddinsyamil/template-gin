package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendOkResponse(c *gin.Context, message string, data interface{}) {
	sendJsonResponse(c, http.StatusOK, message, data)
}

func SendCreatedResponse(c *gin.Context, message string, data interface{}) {
	sendJsonResponse(c, http.StatusCreated, message, data)
}

func SendInternalErrorResponse(c *gin.Context, message string, data interface{}) {
	sendJsonResponse(c, http.StatusInternalServerError, message, data)
}

func SendBadRequestResponse(c *gin.Context, message string, data interface{}) {
	sendJsonResponse(c, http.StatusBadRequest, message, data)
}

func sendJsonResponse(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, map[string]interface{}{
		"message": message,
		"data":    data,
	})
}
