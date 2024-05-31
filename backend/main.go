package main

import (
	"github.com/bagasa11/banturiset/config"

	"github.com/bagasa11/banturiset/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env.local")
	if err != nil {
		panic(err)
	}

	config.InitDB()

	r := gin.Default()
	routes.RegisterRoutes(r)
}
