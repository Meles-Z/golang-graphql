package dbutils

import "gorm.io/gorm"

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate()
	if err != nil {
		return err
	}
	return nil
}
