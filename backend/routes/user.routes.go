package routes

// backend\api\controllers
import (
	"github.com/bagasa11/banturiset/api/controllers"
	"github.com/bagasa11/banturiset/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {
	uc := controllers.NewUsersController()
	r.POST("/register", uc.UserRegistration)
	// create donatur data
	r.POST("/user/donatur/create/:id", uc.DonaturCreate)
	r.POST("/user/peneliti/create/:id", uc.PenelitiCreate)

	r.GET("/user/verify", middleware.JwtAuth(), uc.NeedVerify)
}
