package utils

import (
	"github.com/gin-gonic/gin"
	"messaggio/pkg/model"
)

func NewErrorResponse(c *gin.Context, statusCode int, errorMessage string) {
	c.AbortWithStatusJSON(statusCode, model.ErrorResponse{ErrorMessage: errorMessage})
}
