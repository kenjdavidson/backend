package models

import "github.com/google/uuid"

type Channel struct {
	ChannelID   TwitchID
	ChannelName string
	OverlayID   uuid.UUID
}
