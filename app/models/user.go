package models

type User struct {
	Model
	Name  string `json:"name" gorm:"not null"`
	Email string `json:"email" gorm:"not null;unique"`
}
