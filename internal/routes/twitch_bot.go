package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/streampets/backend/internal/controllers"
)

func RegisterTwitchBotRoutes(r *gin.Engine) {
	api := r.Group("/bot")
	{
		api.GET("/", controllers.HandleIndex)
		api.POST("/join", controllers.HandleJoin)
		api.DELETE("/part", controllers.HandlePart)
		api.PUT("/color", controllers.HandleColor)
		api.POST("/action", controllers.HandleAction)
	}
}
