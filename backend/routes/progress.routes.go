package routes

import (
	"github.com/bagasa11/banturiset/api/controllers"
	"github.com/bagasa11/banturiset/middleware"
	"github.com/gin-gonic/gin"
)

func ProgressRoutes(r *gin.RouterGroup) {
	pc := controllers.NewProgressControllers()
	// r.POST("/project:id/laporan/create", middleware.JwtAuth(), middleware.PenelitiOnly(), middleware.EnsureProjectWasClosed(), middleware.DateAndStageValidate())
	r.POST("/project/:id/laporan-raw/create", middleware.JwtAuth(), middleware.PenelitiOnly(), middleware.IsInState(), pc.Create)
	r.POST("/project/:id/laporan-clean/create", middleware.JwtAuth(), middleware.PenelitiOnly(), middleware.ValidateCreateReport(), pc.Create)

}
