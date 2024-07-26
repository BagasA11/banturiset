package routes

import (
	"net/http"
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
	config.AllowFiles = true
	config.OptionsResponseStatusCode = http.StatusOK
	config.MaxAge = 4 * time.Hour

	// r.Use()
	// r.Use()
	// r.Use()
	useMiddleware := []gin.HandlerFunc{cors.New(config), gin.Logger(), gin.Recovery()}
	apiGroup.Use(useMiddleware...)
	// apiGroup.Use(cors.New(config), gin.Logger(), gin.Recovery())

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
