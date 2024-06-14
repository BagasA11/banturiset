package routes

import (
	"github.com/bagasa11/banturiset/api/controllers"
	"github.com/bagasa11/banturiset/middleware"
	"github.com/gin-gonic/gin"
)

func ProjectRoutes(r *gin.RouterGroup) {
	pc := controllers.NewProjectControllers()
	// peneliti
	r.POST("/project/create", middleware.JwtAuth(), middleware.PenelitiOnly(), pc.Create)
	r.GET("/project/myproject", middleware.JwtAuth(), middleware.PenelitiOnly(), pc.MyProject)
	// umum
	r.GET("/project/:id/detail", middleware.JwtAuth(), pc.Detail)
	// peneliti
	r.PUT("/project/:id/upload/proposal", middleware.JwtAuth(), middleware.PenelitiOnly(), pc.UploadProposal)
	r.PUT("/project/:id/upload/klirens", middleware.JwtAuth(), middleware.PenelitiOnly(), pc.UploadKlirens)
	// admin
	r.PUT("/project/:id/reject", middleware.JwtAuth(), middleware.AdminOnly(), pc.Reject)
}
