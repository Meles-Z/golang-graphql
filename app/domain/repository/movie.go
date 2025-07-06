package repository

import "github.com/meles-z/golang-graphql/app/models"

type MovieRepository interface {
	CreateMovie(*models.Movie) (*models.Movie, error)
	GetMovies() []*models.Movie
	GetMovieById(id string) (*models.Movie, error)
	UpdateMovie(*models.Movie) (*models.Movie, error)
	DeleteMovies(id string) error
}
