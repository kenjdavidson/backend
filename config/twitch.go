package config

import (
	"os"

	"github.com/streampets/backend/repositories"
)

func CreateTwitchRepo() (*repositories.TwitchRepository, error) {
	// Add to GitHub secrets
	// Inject into Docker image
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")

	return repositories.NewTwitchRepository(clientID, clientSecret)
}
