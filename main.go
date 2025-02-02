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

	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}

	twitch, err := config.CreateTwitchRepo()
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	routes.RegisterOverlayRoutes(r, db, twitch)

	r.Run()
}
