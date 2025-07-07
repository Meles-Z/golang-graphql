package db

import (
	"fmt"
	"log"

	"github.com/meles-z/golang-graphql/app/models"
)

// AutoMigrate performs schema migration for all models.
func AutoMigrate() error {
	if DB == nil {
		return fmt.Errorf("database not initialized")
	}

	err := DB.AutoMigrate(
		&models.User{},
		&models.Movie{},
	)

	if err != nil {
		return fmt.Errorf("auto migration failed: %w", err)
	}

	log.Println("✅ Database migration completed")
	return nil
}
