package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/streampets/backend/models"
	"github.com/streampets/backend/repositories"
)

type AuthService struct {
	channelRepo *repositories.ChannelRepository
}

func NewAuthService(channelRepo *repositories.ChannelRepository) *AuthService {
	return &AuthService{channelRepo: channelRepo}
}

func (s *AuthService) VerifyOverlayID(channelID models.UserID, overlayID uuid.UUID) error {
	expectedID, err := s.channelRepo.GetOverlayID(channelID)
	if err != nil {
		return err
	}

	if overlayID != expectedID {
		return errors.New("channelID and overlayID do not match")
	}

	return nil
}
