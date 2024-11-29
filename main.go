package main

import (
	"github.com/gin-gonic/gin"
	"github.com/streampets/backend/config"
	"github.com/streampets/backend/routes"
)

func main() {
	config.ConnectDB()

	r := gin.Default()

	routes.RegisterTwitchBotRoutes(r)

	r.Run()
}
