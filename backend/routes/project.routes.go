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
	r.GET("/project/revisi", middleware.JwtAuth(), middleware.PenelitiOnly(), pc.Revisi)
	r.GET("/project/:id/preview", middleware.JwtAuth(), middleware.PenelitiOnly(), pc.Preview)

	r.PUT("/project/:id/edit", middleware.JwtAuth(), middleware.PenelitiOnly(), pc.Update)
	r.PUT("/project/:id/upload/proposal", middleware.JwtAuth(), middleware.PenelitiOnly(), pc.UploadProposal)
	r.PUT("/project/:id/upload/klirens", middleware.JwtAuth(), middleware.PenelitiOnly(), pc.UploadKlirens)
	r.PUT("/project/:id/submit", middleware.JwtAuth(), middleware.PenelitiOnly(), pc.Submit)

	// admin
	r.GET("/project/hassubmit", middleware.JwtAuth(), middleware.AdminOnly(), pc.HasSubmit)
	r.GET("/project/:id/review", middleware.JwtAuth(), middleware.AdminOnly(), pc.Review)
	r.PUT("/project/:id/reject", middleware.JwtAuth(), middleware.AdminOnly(), pc.Reject)
	r.PUT("/project/:id/verifikasi", middleware.JwtAuth(), middleware.AdminOnly(), pc.Verfikasi)

	// umum
	r.GET("/project/opendonasi", middleware.JwtAuth(), pc.OpenDonate)
	r.GET("/project/diverifikasi", middleware.JwtAuth(), pc.Diverifikasi)
	r.GET("/project/ongoing", middleware.JwtAuth(), pc.OnGoing)
	r.GET("/project/:id/detail", middleware.JwtAuth(), pc.Detail)
	// peneliti

	// admin

}
