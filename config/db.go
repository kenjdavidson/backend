package config

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/streampets/backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	// In Docker compose environment
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	sslMode := os.Getenv("DB_SSL_MODE")
	dbName := os.Getenv("DB_NAME")

	// Convert to Docker secrets
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbName, port, sslMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(
		&models.ChannelItem{},
		&models.Channel{},
		&models.DefaultChannelItem{},
		&models.Item{},
		&models.OwnedItem{},
		&models.Schedule{},
		&models.SelectedItem{},
		&models.User{},
	); err != nil {
		return nil, err
	}

	return db, nil
}
