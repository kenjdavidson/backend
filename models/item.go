package models

import "github.com/google/uuid"

type ItemID uuid.UUID
type Rarity string

const (
	Common   Rarity = "common"
	Uncommon Rarity = "uncommon"
)

type Item struct {
	ItemID  int
	Name    string
	Image   string
	PrevImg string
	Rarity  Rarity
}
