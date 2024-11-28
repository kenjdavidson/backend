package main

import (
	"github.com/gin-gonic/gin"
	"github.com/streampets/backend/internal/routes"
)

func main() {
	r := gin.Default()

	routes.RegisterTwitchBotRoutes(r)

	r.Run()
}
