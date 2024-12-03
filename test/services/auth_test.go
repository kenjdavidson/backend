package services_test

import (
	"slices"
	"testing"

	"github.com/google/uuid"
	"github.com/streampets/backend/models"
	"github.com/streampets/backend/services"
)

type GetOverlayIDCall struct {
	ChannelID models.UserID
}

type SpyRepo struct {
	Calls     []GetOverlayIDCall
	OverlayID uuid.UUID
	Error     error
}

func (spy *SpyRepo) GetOverlayID(channelID models.UserID) (uuid.UUID, error) {
	spy.Calls = append(spy.Calls, GetOverlayIDCall{channelID})
	return spy.OverlayID, spy.Error
}

func TestVerifyOverlayID(t *testing.T) {
	t.Run("verify overlay id returns nil in normal run", func(t *testing.T) {
		channelID := models.UserID("channel id")
		overlayID := uuid.New()

		spyRepo := &SpyRepo{[]GetOverlayIDCall{}, overlayID, nil}

		authService := services.NewAuthService(spyRepo)

		err := authService.VerifyOverlayID(channelID, overlayID)
		if err != nil {
			t.Errorf("did not expect an error, but received one")
		}

		wantedCalls := []GetOverlayIDCall{{channelID}}
		if !slices.Equal(spyRepo.Calls, wantedCalls) {
			t.Errorf("expected %s got %s", wantedCalls, spyRepo.Calls)
		}
	})

	t.Run("verify overlay id returns an error when ids do not match", func(t *testing.T) {
		channelID := models.UserID("channel id")
		overlayID := uuid.New()

		spyRepo := &SpyRepo{[]GetOverlayIDCall{}, uuid.New(), nil}

		authService := services.NewAuthService(spyRepo)

		err := authService.VerifyOverlayID(channelID, overlayID)
		if err == nil {
			t.Errorf("expected an error, but did not received one")
		}

		if err != services.ErrIdMismatch {
			t.Errorf("expected %s got %s", err.Error(), services.ErrIdMismatch.Error())
		}
	})
}
