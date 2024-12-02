package repositories

import (
	"testing"

	"github.com/streampets/backend/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&models.Channel{})
	if err != nil {
		panic(err)
	}

	return db
}

func TestCreateChannel(t *testing.T) {
	t.Run("successful in normal operation", func(t *testing.T) {
		db := setupTestDB()
		repo := NewChannelRepository(db)

		channelID := "test channel id"
		channelName := "test channel name"

		_, err := repo.CreateChannel(channelID, channelName)
		assertNoError(t, err)
	})
}

func assertNoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("did not expect an error, but received one")
	}
}
