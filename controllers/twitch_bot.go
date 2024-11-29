package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type joinParams struct {
	ChannelName string `json:"channel_name" binding:"required"`
	UserID      string `json:"user_id" binding:"required"`
	Username    string `json:"username" binding:"required"`
}

type partParams struct {
	ChannelName string `json:"channel_name" binding:"required"`
	UserID      string `json:"user_id" binding:"required"`
}

type colorParams struct {
	ChannelName string `json:"channel_name" binding:"required"`
	UserID      string `json:"user_id" binding:"required"`
	Color       string `json:"color" binding:"required"`
}

type actionParams struct {
	ChannelName string `json:"channel_name" binding:"required"`
	UserID      string `json:"user_id" binding:"required"`
	Action      string `json:"action" binding:"required"`
}

func HandleJoin(ctx *gin.Context) {
	var params joinParams

	if err := ctx.ShouldBindJSON(params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Get channel_id from channel_name
	// Get user's params from DB -> Viewer
	// Send Viewer in SSE to overlay
}

func HandlePart(ctx *gin.Context) {
	var params partParams

	if err := ctx.ShouldBindJSON(params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// ? - Get channel_id from channel_name
	// Send UserID in SSE to overlay
}

func HandleColor(ctx *gin.Context) {
	var params colorParams

	if err := ctx.ShouldBindJSON(params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Update the color of UserID & ChannelID in DB
	// Send SSE to ChannelName's overlay
}

func HandleAction(ctx *gin.Context) {
	var params actionParams

	if err := ctx.ShouldBindJSON(params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Send Action&UserID to ChannelName's overlay
}
