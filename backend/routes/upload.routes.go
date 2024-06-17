package routes

import (
	"github.com/bagasa11/banturiset/api/controllers"
	"github.com/bagasa11/banturiset/middleware"
	"github.com/gin-gonic/gin"
)

func UploadRoutes(r *gin.RouterGroup) {
	r.POST("file/upload", middleware.JwtAuth(), controllers.Upload)
	// r.POST("/upload/multi", middleware.JwtAuth(), controllers.UploadMulti)
	r.GET("file/download", middleware.JwtAuth(), controllers.Download)
}
