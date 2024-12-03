package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/streampets/backend/controllers"
	"github.com/streampets/backend/repositories"
	"github.com/streampets/backend/services"
	"gorm.io/gorm"
)

func RegisterOverlayRoutes(r *gin.Engine, db *gorm.DB, twitchRepo *repositories.TwitchRepository) {
	channelRepo := repositories.NewChannelRepository(db)

	authService := services.NewAuthService(channelRepo)
	channelService := services.NewChannelService(twitchRepo)

	overlayController := controllers.NewOverlayController(channelService, authService)

	api := r.Group("/api/v1")
	{
		api.GET("/listen", overlayController.HandleListen)
		api.GET("/viewers", overlayController.HandleViewers)
	}
}
