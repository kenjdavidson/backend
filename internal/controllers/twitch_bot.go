package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/streampets/backend/internal/models"
)

func HandleIndex(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}

func HandleJoin(ctx *gin.Context) {
	var data models.Join

	if err := ctx.ShouldBindJSON(data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Get channel_id from channel_name
	// Get user's data from DB -> Viewer
	// Send Viewer in SSE to overlay
}

func HandlePart(ctx *gin.Context) {
	var data models.Part

	if err := ctx.ShouldBindJSON(data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// ? - Get channel_id from channel_name
	// Send UserID in SSE to overlay
}

func HandleColor(ctx *gin.Context) {
	var data models.Color

	if err := ctx.ShouldBindJSON(data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Update the color of UserID & ChannelID in DB
	// Send SSE to ChannelName's overlay
}

func HandleAction(ctx *gin.Context) {
	var data models.Action

	if err := ctx.ShouldBindJSON(data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Send Action&UserID to ChannelName's overlay
}
