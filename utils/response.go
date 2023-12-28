package utils

import (
	"github.com/gin-gonic/gin"
)

func SendError(ctx *gin.Context, code int, err error) {
	ctx.JSON(code, gin.H{"message": err.Error()})
}

func SendSuccess(ctx *gin.Context, code int, data any) {
	ctx.JSON(code, gin.H{
		"messsage": "success",
		"data":     data,
	})
}
