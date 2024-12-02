package models

import "github.com/google/uuid"

type OwnedItem struct {
	UserID        string
	ChannelID     string
	ItemID        uuid.UUID
	TransactionID uuid.UUID
}
