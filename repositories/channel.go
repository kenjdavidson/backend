package repositories

import (
	"github.com/google/uuid"
	"github.com/streampets/backend/models"
	"gorm.io/gorm"
)

type ChannelRepository struct {
	db *gorm.DB
}

func NewChannelRepository(db *gorm.DB) *ChannelRepository {
	return &ChannelRepository{db: db}
}

func (repo *ChannelRepository) CreateChannel(channelID, channelName string) (uuid.UUID, error) {
	overlayID := uuid.New()

	channel := models.Channel{ChannelID: channelID, ChannelName: channelName, OverlayID: overlayID}
	result := repo.db.Create(&channel)

	return overlayID, result.Error
}
