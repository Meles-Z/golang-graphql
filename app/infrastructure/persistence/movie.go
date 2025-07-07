package persistence

import (
	"context"
	"fmt"

	"github.com/meles-z/golang-graphql/app/domain/repository"
	"github.com/meles-z/golang-graphql/app/models"
	"gorm.io/gorm"
)

type movieRepositoryImpl struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) repository.MovieRepository {
	return &movieRepositoryImpl{db}
}

func (r *movieRepositoryImpl) Create(ctx context.Context, input models.MovieInput) (*models.Movie, error) {
	movie := models.Movie{
		Title:       input.Title,
		URL:         input.URL,
		ReleaseDate: input.ReleaseDate,
	}
	if err := r.db.WithContext(ctx).Create(&movie).Error; err != nil {
		return nil, err
	}
	return &movie, nil
}

func (r *movieRepositoryImpl) Update(ctx context.Context, id string, input models.MovieInput) (*models.Movie, error) {
	var movie models.Movie
	if err := r.db.WithContext(ctx).First(&movie, "id = ?", id).Error; err != nil {
		return nil, err
	}
	movie.Title = input.Title
	movie.URL = input.URL
	movie.ReleaseDate = input.ReleaseDate
	if err := r.db.WithContext(ctx).Save(&movie).Error; err != nil {
		return nil, err
	}
	return &movie, nil
}

func (r *movieRepositoryImpl) Delete(ctx context.Context, id string) (bool, error) {
	movie, err := r.GetByID(ctx, id)
	if err != nil {
		return false, fmt.Errorf("user not found")
	}
	if err := r.db.WithContext(ctx).Delete(movie).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *movieRepositoryImpl) GetByID(ctx context.Context, id string) (*models.Movie, error) {
	var movie models.Movie
	if err := r.db.WithContext(ctx).First(&movie, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &movie, nil
}

func (r *movieRepositoryImpl) GetAll(ctx context.Context, filter *models.MovieFilter) ([]*models.Movie, error) {
	var movies []*models.Movie
	query := r.db.WithContext(ctx).Model(&models.Movie{})
	// Apply filter if needed
	if filter != nil && filter.Title != nil {
		query = query.Where("title ILIKE ?", "%"+*filter.Title+"%")
	}
	if err := query.Find(&movies).Error; err != nil {
		return nil, err
	}
	return movies, nil
}
