package routes

import (
	"github.com/bagasa11/banturiset/api/controllers"
	"github.com/bagasa11/banturiset/middleware"
	"github.com/gin-gonic/gin"
)

func BudgetRoutes(r *gin.RouterGroup) {
	bc := controllers.NewBudgetDetailsController()
	r.POST("/project/:id/budget/create", middleware.JwtAuth(), middleware.PenelitiOnly(), bc.Create)
	r.PUT("/project/:id/budget/:budgetid/update", middleware.JwtAuth(), middleware.PenelitiOnly(), bc.Updates)
	r.DELETE("/project/:id/budget/:budgetid/delete", middleware.JwtAuth(), middleware.PenelitiOnly(), bc.Delete)
}
