package routes

import (
	"os"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	apiGroup := r.Group("/api")
	UserRoutes(apiGroup)

	r.Run(":" + os.Getenv("LOC_PORT"))
}
