package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/streampets/backend/repositories"
)

type event struct {
	Event   string
	Message string
}

type OverlayService struct {
	repo  *repositories.OverlayRepo
	chans map[string]chan event
}

func NewOverlayService(repo *repositories.OverlayRepo) *OverlayService {
	return &OverlayService{repo: repo}
}

func (s *OverlayService) VerifyOverlayID(channelID string, overlayID uuid.UUID) error {
	expectedID, err := s.repo.GetOverlayID(channelID)
	if err != nil {
		return err
	}

	if overlayID != expectedID {
		return errors.New("channelID and overlayID do not match")
	}

	return nil
}

func (s *OverlayService) GetChannel(channelID string) (chan event, error) {
	val, ok := s.chans[channelID]
	if ok {
		return val, nil
	}

	return nil, errors.New("channel id does not have an event stream")
}
