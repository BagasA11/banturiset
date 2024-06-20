package routes

import (
	"github.com/bagasa11/banturiset/api/controllers"
	"github.com/bagasa11/banturiset/middleware"
	"github.com/gin-gonic/gin"
)

func DonasiRoutes(r *gin.RouterGroup) {
	dc := controllers.NewDonasi()

	r.POST("/project/:id/donasi/create", middleware.JwtAuth(), middleware.DonaturOnly(), dc.CreateDonasi)
	r.GET("/donasi/:id", middleware.JwtAuth(), middleware.DonaturOnly(), dc.GetInvoiceDetail)
	r.GET("/project/:id/donasi/histori", middleware.JwtAuth(), dc.GetHistory)
	r.GET("/project/:id/donasi/contributor", middleware.JwtAuth(), dc.Contributors)
	r.GET("/donasi/mycontribution", middleware.JwtAuth(), middleware.DonaturOnly(), dc.Contributors)
	r.POST("/donasi/notif", middleware.VerifyWebhookToken(), middleware.CheckRedundantWebhook(), dc.Notif)
}
