package routes

import (
	"github.com/bagasa11/banturiset/api/controllers"
	"github.com/bagasa11/banturiset/middleware"
	"github.com/gin-gonic/gin"
)

func TahapRoutes(r *gin.RouterGroup) {
	tc := controllers.NewTahapControllers()
	r.POST("/project/:id/tahap/create", middleware.JwtAuth(), tc.Create)
	r.GET("/project/:id/tahap/list", middleware.JwtAuth(), tc.List)
	r.PUT("/project/tahap/update/:id", middleware.JwtAuth(), tc.Update)
	r.DELETE("/project/tahap/delete/:id", middleware.JwtAuth(), tc.Delete)
}
