package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/streampets/backend/controllers"
)

func RegisterTwitchBotRoutes(r *gin.Engine) {
	api := r.Group("/bot")
	{
		api.POST("/join", controllers.HandleJoin)
		api.DELETE("/part", controllers.HandlePart)
		api.PUT("/color", controllers.HandleColor)
		api.POST("/action", controllers.HandleAction)
	}
}
