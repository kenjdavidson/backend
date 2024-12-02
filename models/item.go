package models

import "github.com/google/uuid"

type Rarity string

const (
	Common   Rarity = "common"
	Uncommon Rarity = "uncommon"
)

type Item struct {
	ItemID  uuid.UUID
	Name    string
	Rarity  Rarity
	Image   string
	PrevImg string
}
