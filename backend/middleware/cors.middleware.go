package middleware

import (
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("Access-Control-Allow-Origin", "*")
		ctx.Next()
	}
}
