package services

import (
	"github.com/streampets/backend/models"
)

type Event struct {
	Event   string
	Message string
}

type EventStream chan Event

type TwitchRepo interface {
	GetUsername(channelID models.UserID) (string, error)
}

type ChannelService struct {
	twitchRepo TwitchRepo
	streams    map[string]EventStream
}

func NewChannelService(twitchRepo TwitchRepo) *ChannelService {
	return &ChannelService{
		twitchRepo: twitchRepo,
		streams:    make(map[string]EventStream),
	}
}

func (s *ChannelService) GetEventStream(channelID models.UserID) (EventStream, error) {
	channelName, err := s.twitchRepo.GetUsername(channelID)
	if err != nil {
		return nil, err
	}

	stream, ok := s.streams[channelName]
	if ok {
		return stream, nil
	}

	stream = make(EventStream)
	s.streams[channelName] = stream

	return stream, nil
}
