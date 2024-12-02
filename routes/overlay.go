package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/streampets/backend/controllers"
	"github.com/streampets/backend/repositories"
	"github.com/streampets/backend/services"
	"gorm.io/gorm"
)

func RegisterOverlayRoutes(r *gin.Engine, db *gorm.DB) {
	overlayRepo := repositories.NewOverlayRepo(db)
	overlayService := services.NewOverlayService(overlayRepo)
	overlayController := controllers.NewOverlayController(overlayService)

	api := r.Group("/api/v1")
	{
		api.GET("/listen", overlayController.HandleListen)
		api.GET("/viewers", overlayController.HandleViewers)
	}
}
