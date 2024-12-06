package services

import (
	"github.com/streampets/backend/models"
)

type Event struct {
	Event   string
	Message string
}

type EventStream chan Event

type Viewer struct {
}

type TwitchRepo interface {
	GetUsername(channelID models.TwitchID) (string, error)
}

type ChannelService struct {
	twitchRepo TwitchRepo
	streams    map[string]EventStream
	viewers    map[models.TwitchID]([]Viewer)
}

func NewChannelService(twitchRepo TwitchRepo) *ChannelService {
	return &ChannelService{
		twitchRepo: twitchRepo,
		streams:    make(map[string]EventStream),
		viewers:    make(map[models.TwitchID]([]Viewer)),
	}
}

func (s *ChannelService) GetEventStream(channelID models.TwitchID) (EventStream, error) {
	channelName, err := s.twitchRepo.GetUsername(channelID)
	if err != nil {
		return nil, err
	}

	stream, ok := s.streams[channelName]
	if !ok {
		stream = make(EventStream)
		s.streams[channelName] = stream
	}

	return stream, nil
}

func (s *ChannelService) GetChannelsViewers(channelID models.TwitchID) ([]Viewer, error) {
	viewers, ok := s.viewers[channelID]
	if !ok {
		viewers = []Viewer{}
		s.viewers[channelID] = viewers
	}

	return viewers, nil
}
