package repositories_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/streampets/backend/models"
	"github.com/streampets/backend/repositories"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err = db.AutoMigrate(&models.Channel{}); err != nil {
		panic(err)
	}

	return db
}

func TestGetOverlayID(t *testing.T) {
	db := setupTestDB()
	repo := repositories.NewChannelRepository(db)

	channelID := models.TwitchID("channel id")
	channelName := "channel name"
	expectedID := uuid.New()

	channel := models.Channel{ChannelID: channelID, ChannelName: channelName, OverlayID: expectedID}

	if err := db.Create(&channel).Error; err != nil {
		t.Errorf("could not create record")
	}

	overlayID, err := repo.GetOverlayID(channelID)
	if err != nil {
		t.Errorf("did not expect an error but received: %s", err.Error())
	}

	if overlayID != expectedID {
		t.Errorf("expected %s got %s", expectedID, overlayID)
	}
}
