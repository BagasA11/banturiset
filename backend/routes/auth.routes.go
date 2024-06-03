package routes

import (
	"github.com/bagasa11/banturiset/api/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup) {
	ac := controllers.NewAuthController()
	r.POST("/login", ac.Login)
	r.POST("/refresh", ac.RefreshToken)
}
