package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/streampets/backend/config"
	"github.com/streampets/backend/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	config.ConnectDB()

	r := gin.Default()

	routes.RegisterTwitchBotRoutes(r)

	r.Run()
}
