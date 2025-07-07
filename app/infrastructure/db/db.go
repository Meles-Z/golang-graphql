package db

import (
	"fmt"
	"log"

	"github.com/meles-z/golang-graphql/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes and returns a gorm.DB instance.
func InitDB(cfg configs.DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.DBName, cfg.User, cfg.Password,
	)

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("✅ Connected to database successfully")
	DB = conn
	return conn, nil
}
