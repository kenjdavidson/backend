package services_test

import (
	"errors"
	"testing"

	"github.com/streampets/backend/models"
	"github.com/streampets/backend/services"
)

type GetUsernameCall struct {
	ChannelID models.UserID
}

type spyTwitchRepo struct {
	Calls     []GetUsernameCall
	Usernames map[models.UserID]string
	Error     error
}

func (spy *spyTwitchRepo) GetUsername(channelID models.UserID) (string, error) {
	spy.Calls = append(spy.Calls, GetUsernameCall{channelID})
	val, ok := spy.Usernames[channelID]
	if ok {
		return val, nil
	}
	return "", errors.New("invalid channelid")
}

func TestGetEventStream(t *testing.T) {
	t.Run("return the same stream for the same channel id", func(t *testing.T) {
		channelID := models.UserID("channel id")

		spyTwitchRepo := &spyTwitchRepo{
			[]GetUsernameCall{},
			map[models.UserID]string{
				channelID: "username",
			},
			nil,
		}

		channelService := services.NewChannelService(spyTwitchRepo)

		streamOne, err := channelService.GetEventStream(channelID)
		if err != nil {
			t.Errorf("did not expect an error, but received one")
		}

		streamTwo, err := channelService.GetEventStream(channelID)
		if err != nil {
			t.Errorf("did not expect an error, but received one")
		}

		if streamOne != streamTwo {
			t.Errorf("the two streams are not the same")
		}
	})

	t.Run("return different streams for different channel ids", func(t *testing.T) {
		channelIdOne := models.UserID("channel id one")
		channelIdTwo := models.UserID("channel id two")

		spyTwitchRepo := &spyTwitchRepo{
			[]GetUsernameCall{},
			map[models.UserID]string{
				channelIdOne: "username one",
				channelIdTwo: "username two",
			},
			nil,
		}

		channelService := services.NewChannelService(spyTwitchRepo)

		streamOne, err := channelService.GetEventStream(channelIdOne)
		if err != nil {
			t.Errorf("did not expect an error, but received one")
		}

		streamTwo, err := channelService.GetEventStream(channelIdTwo)
		if err != nil {
			t.Errorf("did not expect an error, but received one")
		}

		if streamOne == streamTwo {
			t.Errorf("the two streams are the same")
		}
	})
}
