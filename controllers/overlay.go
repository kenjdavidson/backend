package controllers

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/streampets/backend/services"
)

type listenParams struct {
	ChannelID string
	OverlayID uuid.UUID
}

type viewerParams struct {
	ChannelID string
	OverlayID uuid.UUID
}

type OverlayController struct {
	service *services.OverlayService
}

func NewOverlayController(service *services.OverlayService) *OverlayController {
	return &OverlayController{service: service}
}

func (c *OverlayController) HandleListen(ctx *gin.Context) {
	var params listenParams

	if err := ctx.ShouldBindJSON(params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := c.service.VerifyOverlayID(params.ChannelID, params.OverlayID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ch, err := c.service.GetChannel(params.ChannelID)
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

func (c *OverlayController) HandleViewers(ctx *gin.Context) {
	var params viewerParams

	if err := ctx.ShouldBindJSON(params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

}
