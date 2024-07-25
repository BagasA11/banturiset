package routes

import (
	"os"

	"github.com/bagasa11/banturiset/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	apiGroup := r.Group("/api")
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// cors
	r.Use(middleware.DefaultConfigCors())
	UserRoutes(apiGroup)
	AuthRoutes(apiGroup)
	UploadRoutes(apiGroup)
	PengajuanRoutes(apiGroup)
	ProjectRoutes(apiGroup)
	TahapRoutes(apiGroup)
	BudgetRoutes(apiGroup)
	DonasiRoutes(apiGroup)
	ProgressRoutes(apiGroup)
	r.Run(":" + os.Getenv("LOC_PORT"))
}
