package models

type TwitchID string

type User struct {
	UserID   TwitchID
	Username string
}
