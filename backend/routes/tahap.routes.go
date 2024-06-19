package routes

import (
	"github.com/bagasa11/banturiset/api/controllers"
	"github.com/bagasa11/banturiset/middleware"
	"github.com/gin-gonic/gin"
)

func TahapRoutes(r *gin.RouterGroup) {
	tc := controllers.NewTahapControllers()
	r.POST("/project/:id/tahap/create", middleware.JwtAuth(), middleware.PenelitiOnly(), tc.Create)
	r.GET("/project/:id/tahap/list", middleware.JwtAuth(), tc.List)
	r.PUT("/project/:id/tahap/update/:tahapid", middleware.JwtAuth(), middleware.PenelitiOnly(), tc.Update)
	r.DELETE("/project/:id/tahap/delete/:tahapid", middleware.JwtAuth(), middleware.PenelitiOnly(), tc.Delete)
}
