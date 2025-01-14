package routes

import (
	"net/http"
	"os"
	"time"

	mw "github.com/bagasa11/banturiset/middleware"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Access-Control-Allow-Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.AllowFiles = true
	config.OptionsResponseStatusCode = http.StatusOK
	config.MaxAge = 4 * time.Hour

	useMiddleware := []gin.HandlerFunc{cors.New(config),
		gin.Logger(), gin.Recovery(),
		mw.SimpleLimiter(),
	}
	r.Use(useMiddleware...)

	apiGroup := r.Group("/api")
	UserRoutes(apiGroup)
	AuthRoutes(apiGroup)
	UploadRoutes(apiGroup)
	PengajuanRoutes(apiGroup)
	ProjectRoutes(apiGroup)
	BudgetRoutes(apiGroup)
	DonasiRoutes(apiGroup)
	// ProgressRoutes(apiGroup)
	// staticDir := "./../file/gambar"
	r.StaticFS("/public", gin.Dir("file", true))
	r.Run(":" + os.Getenv("LOC_PORT"))

}
