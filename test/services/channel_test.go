package services_test

import (
	"errors"
	"slices"
	"testing"

	"github.com/streampets/backend/models"
	"github.com/streampets/backend/repositories"
	"github.com/streampets/backend/services"
)

type GetUsernameCall struct {
	ChannelID models.TwitchID
}

type spyTwitchRepo struct {
	Calls     []GetUsernameCall
	Usernames map[models.TwitchID]string
	Error     error
}

func (spy *spyTwitchRepo) GetUsername(channelID models.TwitchID) (string, error) {
	spy.Calls = append(spy.Calls, GetUsernameCall{channelID})
	val, ok := spy.Usernames[channelID]
	if ok {
		return val, nil
	}
	return "", errors.New("invalid channelid")
}

func TestGetEventStream(t *testing.T) {
	t.Run("return the same stream for the same channel id", func(t *testing.T) {
		channelID := models.TwitchID("channel id")

		spyTwitchRepo := &spyTwitchRepo{
			[]GetUsernameCall{},
			map[models.TwitchID]string{
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
		channelIdOne := models.TwitchID("channel id one")
		channelIdTwo := models.TwitchID("channel id two")

		spyTwitchRepo := &spyTwitchRepo{
			[]GetUsernameCall{},
			map[models.TwitchID]string{
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

func TestGetChannelsViewers(t *testing.T) {
	channelID := models.TwitchID("channel id")

	twitchRepo := &repositories.TwitchRepository{}
	channelService := services.NewChannelService(twitchRepo)

	viewers, err := channelService.GetChannelsViewers(channelID)
	if err != nil {
		t.Errorf("did not expect an error but received %s", err.Error())
	}

	expected := []services.Viewer{}
	if !slices.Equal(viewers, expected) {
		t.Errorf("expected %s got %s", expected, viewers)
	}
}
