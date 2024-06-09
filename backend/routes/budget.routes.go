package routes

import (
	"github.com/bagasa11/banturiset/api/controllers"
	"github.com/bagasa11/banturiset/middleware"
	"github.com/gin-gonic/gin"
)

func BudgetRoutes(r *gin.RouterGroup) {
	bc := controllers.NewBudgetDetailsController()
	r.POST("/project/:id/budget/create", middleware.JwtAuth(), bc.Create)
	r.PUT("/project/budget/:id/update", middleware.JwtAuth(), bc.Updates)
	r.DELETE("/project/budget/:id/delete", middleware.JwtAuth(), bc.Delete)
}
