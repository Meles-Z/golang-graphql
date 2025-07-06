package db

import (
	"fmt"

	"gorm.io/gorm"
)


func InitDD() (*gorm.DB, error){
	dsn:=fmt.Sprintf("host=%s port=%d dbname=%s user=%s password+%s")
}