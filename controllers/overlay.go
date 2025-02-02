package controllers

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/streampets/backend/models"
	"github.com/streampets/backend/services"
)

type chanService interface {
	GetEventStream(channelID models.TwitchID) (services.EventStream, error)
	GetChannelsViewers(channelID models.TwitchID) ([]services.Viewer, error)
}

type authService interface {
	VerifyOverlayID(models.TwitchID, uuid.UUID) error
}

type OverlayController struct {
	chanService chanService
	authService authService
}

func NewOverlayController(chanService chanService, authService authService) *OverlayController {
	return &OverlayController{
		chanService: chanService,
		authService: authService,
	}
}

func (c *OverlayController) HandleListen(ctx *gin.Context) {
	channelID := models.TwitchID(ctx.Query("channelID"))
	overlayID, err := uuid.Parse(ctx.Query("overlayID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := c.authService.VerifyOverlayID(channelID, overlayID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	viewers, err := c.chanService.GetChannelsViewers(channelID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	for viewer := range viewers {
		// TODO: Ensure viewer is formatted correctly
		ctx.SSEvent("JOIN", viewer)
	}

	ch, err := c.chanService.GetEventStream(channelID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.Stream(func(w io.Writer) bool {
		if event, ok := <-ch; ok {
			ctx.SSEvent(event.Event, event.Message)
			return true
		}
		return false
	})
}
