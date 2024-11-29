package models

type UserID string

type User struct {
	UserID       UserID
	Username     string
	OwnedItems   map[ChannelID]([]Item)
	SelectedItem map[ChannelID](Item)
}
