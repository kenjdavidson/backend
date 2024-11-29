package models

type ChannelID string
type DayOfWeek string

const (
	Monday    DayOfWeek = "monday"
	Tuesday   DayOfWeek = "tuesday"
	Wednesday DayOfWeek = "wednesday"
	Thursday  DayOfWeek = "thursday"
	Friday    DayOfWeek = "friday"
	Saturday  DayOfWeek = "saturday"
	Sunday    DayOfWeek = "sunday"
)

type Channel struct {
	ChannelID   string
	ChannelName string
	OverlayID   string
	Schedule    map[DayOfWeek]([]Item)
	DefaultItem Item
	Items       []Item
}
