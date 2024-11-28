package models

type Join struct {
	ChannelName string `json:"channel_name" binding:"required"`
	UserID      string `json:"user_id" binding:"required"`
	Username    string `json:"username" binding:"required"`
}

type Part struct {
	ChannelName string `json:"channel_name" binding:"required"`
	UserID      string `json:"user_id" binding:"required"`
}

type Color struct {
	ChannelName string `json:"channel_name" binding:"required"`
	UserID      string `json:"user_id" binding:"required"`
	Color       string `json:"color" binding:"required"`
}

type Action struct {
	ChannelName string `json:"channel_name" binding:"required"`
	UserID      string `json:"user_id" binding:"required"`
	Action      string `json:"action" binding:"required"`
}
