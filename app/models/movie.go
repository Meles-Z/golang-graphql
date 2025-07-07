package models

type Movie struct {
	Model
	Title       string `json:"title" gorm:"not null"`
	URL         string `json:"url" gorm:"not null;unique"`
	ReleaseDate string `json:"releaseDate" gorm:"not null"`
}

