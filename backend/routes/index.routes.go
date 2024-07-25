package routes

import (
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	apiGroup := r.Group("/api")
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour

	r.Use(cors.New(config))
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

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
