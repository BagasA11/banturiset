package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/autotls"
	"golang.org/x/crypto/acme/autocert"

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

	// r.Use()
	// r.Use()
	// r.Use()
	useMiddleware := []gin.HandlerFunc{cors.New(config), gin.Logger(), gin.Recovery()}
	r.Use(useMiddleware...)

	apiGroup := r.Group("/api")
	// apiGroup.Use(useMiddleware...)
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
	// r.Run(":" + os.Getenv("LOC_PORT"))

	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("bagasa11.my.id"),
		Cache:      autocert.DirCache("/var/www/.cache"),
	}

	autotls.RunWithManager(r, &m)
}
