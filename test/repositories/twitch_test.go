package repositories_test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/streampets/backend/models"
	"github.com/streampets/backend/repositories"
)

func setupTwitchRepository() *repositories.TwitchRepository {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic(err)
	}

	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")

	return repositories.NewTwitchRepository(clientID, clientSecret)
}

func TestGetUsername(t *testing.T) {
	twitchRepo := setupTwitchRepository()

	channelID := models.UserID("83125762")
	expected := "ljrexcodes"

	username, err := twitchRepo.GetUsername(channelID)
	if err != nil {
		t.Errorf("did not expect an error but received: %s", err.Error())
	}

	if username != expected {
		t.Errorf("expected %s got %s", expected, username)
	}
}
