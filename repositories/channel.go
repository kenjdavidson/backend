package repositories

import (
	"database/sql"
)

type ChannelRepository struct {
	DB *sql.DB
}

func NewChannelRepository(db *sql.DB) *ChannelRepository {
	return &ChannelRepository{DB: db}
}

func (repo *ChannelRepository) CreateChannel(channelID, channelName, overlayID string) error {
	tx, err := repo.DB.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}

	const createChannelQuery = `
		INSERT INTO channels (channelid, channelname, overlayid)
		VALUES ($1,$2,$3)`
	_, err = repo.DB.Exec(createChannelQuery, channelID, channelName, overlayID)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
