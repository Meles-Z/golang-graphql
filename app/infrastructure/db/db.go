package db

import (
	"fmt"

	"github.com/meles-z/golang-graphql/app/models"
	"github.com/meles-z/golang-graphql/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(cfg configs.DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s",
		cfg.Host, cfg.Port, cfg.DBName, cfg.User, cfg.Password)
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	fmt.Println("Database connected successfully")
	conn.AutoMigrate(&models.User{}, &models.Movie{})

	conn = DB
	return conn, nil
}
