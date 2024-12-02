package repositories

import (
	"github.com/google/uuid"
	"github.com/streampets/backend/models"
	"gorm.io/gorm"
)

type ChannelRepository struct {
	DB *gorm.DB
}

func NewChannelRepository(db *gorm.DB) *ChannelRepository {
	return &ChannelRepository{DB: db}
}

func (repo *ChannelRepository) CreateChannel(channelID, channelName string) (uuid.UUID, error) {
	overlayID := uuid.New()

	channel := models.Channel{ChannelID: channelID, ChannelName: channelName, OverlayID: overlayID}
	result := repo.DB.Create(&channel)

	return overlayID, result.Error
}
