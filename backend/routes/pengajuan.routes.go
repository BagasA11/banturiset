package routes

import (
	"github.com/bagasa11/banturiset/api/controllers"
	"github.com/bagasa11/banturiset/middleware"
	"github.com/gin-gonic/gin"
)

func PengajuanRoutes(r *gin.RouterGroup) {
	pc := controllers.NewPengajuanController()
	r.POST("/pengajuan/create", middleware.JwtAuth(), pc.Create)
	r.GET("/pengajuan/open", pc.Open)
}
