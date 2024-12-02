package models

import "github.com/google/uuid"

type ChannelItem struct {
	ChannelID string
	ItemID    uuid.UUID
}
