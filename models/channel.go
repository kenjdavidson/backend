package models

import "github.com/google/uuid"

type ChannelName string

type Channel struct {
	ChannelID   UserID
	ChannelName ChannelName
	OverlayID   uuid.UUID
}
