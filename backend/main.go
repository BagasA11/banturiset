package main

import (
	"fmt"
	"os"

	"github.com/bagasa11/banturiset/config"

	"github.com/bagasa11/banturiset/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env.local")
	if err != nil {
		panic(err)
	}

	fmt.Println(os.Getenv("LOC_HOST"))
	if err := config.InitDB(); err != nil {
		panic(err)
	}
	// initialize ttl cache
	config.InitCache()
	r := gin.Default()
	routes.RegisterRoutes(r)
}
