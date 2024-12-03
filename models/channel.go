package models

import "github.com/google/uuid"

type Channel struct {
	ChannelID   UserID
	ChannelName string
	OverlayID   uuid.UUID
}
