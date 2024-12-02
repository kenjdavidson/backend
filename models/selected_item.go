package models

import "github.com/google/uuid"

type SelectedItem struct {
	UserID    string
	ChannelID string
	ItemID    uuid.UUID
}
