package main

import (
	"github.com/bagasa11/banturiset/config"
	"github.com/bagasa11/banturiset/routes"
	"github.com/bagasa11/banturiset/timezone"
	val "github.com/bagasa11/banturiset/validators"

	"github.com/gin-gonic/gin"
	bind "github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func main() {

	// home := "/home/banturiset/backend/"
	err := godotenv.Load(".env.local")
	if err != nil {
		panic(err)
	}

	// initialize timezone
	if err := timezone.SetLocation(timezone.Timezone); err != nil {
		panic(err)
	}

	// fmt.Println(os.Getenv("LOC_HOST"))
	if err := config.InitDB(); err != nil {
		panic(err)
	}

	if v, ok := bind.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("start", val.ValidateStartTime)
	}

	// initialize ttl cache
	config.InitCache()
	// gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	routes.RegisterRoutes(r)
}
