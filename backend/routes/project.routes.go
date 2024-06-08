package routes

import (
	"github.com/bagasa11/banturiset/api/controllers"
	"github.com/bagasa11/banturiset/middleware"
	"github.com/gin-gonic/gin"
)

func ProjectRoutes(r *gin.RouterGroup) {
	pc := controllers.NewProjectControllers()
	r.POST("/project/create", middleware.JwtAuth(), pc.Create)
	r.GET("/project/myproject", middleware.JwtAuth(), pc.MyProject)
	r.PUT("/project/:id/upload/proposal", middleware.JwtAuth(), pc.UploadProposal)
	r.PUT("/project/:id/upload/klirens", middleware.JwtAuth(), pc.UploadProposal)
	r.PUT("/project/:id/reject", middleware.JwtAuth(), pc.Reject)
}
