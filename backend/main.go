package main

import (
	"fmt"

	"github.com/bagasa11/banturiset/config"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env.local")
	if err != nil {
		panic(err)
	}

	fmt.Println("error", config.InitDB())
}
