package config

import (
	"os"

	"github.com/streampets/backend/repositories"
)

func CreateTwitchRepo() *repositories.TwitchRepository {
	// Stored in GitHub secrets
	// Injected into image's env at build
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")

	return repositories.NewTwitchRepository(clientID, clientSecret)
}
