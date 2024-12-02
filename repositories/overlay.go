package repositories

import (
	"github.com/google/uuid"
	"github.com/streampets/backend/models"
	"gorm.io/gorm"
)

type OverlayRepo struct {
	db *gorm.DB
}

func NewOverlayRepo(db *gorm.DB) *OverlayRepo {
	return &OverlayRepo{db: db}
}

func (repo *OverlayRepo) GetOverlayID(channelID string) (uuid.UUID, error) {
	var channel models.Channel

	result := repo.db.Table("channels").Where("channel_id = ?", channelID).First(channel)

	return channel.OverlayID, result.Error
}
